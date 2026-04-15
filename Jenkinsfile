// ── Parameters ────────────────────────────────────────────────────────────────
properties([
    parameters([
        string(name: 'GM_URL',     defaultValue: 'https://host-a.infoblox.com', description: 'URL For the GRID Master'),
        string(name: 'MEMBER_URL', defaultValue: 'https://host-b.infoblox.com', description: 'URL for the GRID Member'),
        string(name: 'TF_VERSION', defaultValue: '1.8.0',                       description: 'Terraform version to install e.g. 1.8.0')
    ])
])

def runTestStage(String name, String dir, String timeout) {
    stage("Tests: ${name}") {
        try {
            sh """
                cd ${dir}
                go test -v -count=1 -timeout ${timeout} ./... \
                    2>&1 | tee \$WORKSPACE/test-results/${name}.txt
            """
        } catch (err) {
            echo "Tests in ${dir} failed: ${err}"
            currentBuild.result = 'FAILURE'
        } finally {
            sh "go-junit-report < \$WORKSPACE/test-results/${name}.txt > \$WORKSPACE/test-results/${name}-junit.xml || true"
        }
    }
}


// ── Main pipeline ─────────────────────────────────────────────────────────────
node('Cloud-test1-172.28.81.12-label') {
    timestamps {
        timeout(time: 240, unit: 'MINUTES') {

            // ── Environment variables ─────────────────────────────────────────
            env.NIOS_HOST_URL     = params.GM_URL
            env.NIOS_MEMBER_URL   = params.MEMBER_URL
            env.NIOS_WAPI_VERSION = 'v2.13.6'
            env.TF_ACC            = '1'
            env.GO_VERSION        = '1.25.1'
            env.TF_VERSION        = params.TF_VERSION
            env.PATH              = "${env.WORKSPACE}/tools/go/bin:/usr/local/go/bin:${env.HOME}/go/bin:${env.WORKSPACE}/bin:${env.PATH}"

            try {

                // ── Checkout ──────────────────────────────────────────────────
                stage('Checkout') {
                    git url: 'https://github.com/unasra/terraform-provider-nios.git', branch: 'fix_tests'
                }

                // ── Install toolchain ─────────────────────────────────────────
                stage('Install toolchain') {
                    sh '''
                        # ── Go ───────────────────────────────────────────────
                        INSTALLED_GO=$(go version 2>/dev/null | awk '{print $3}' | sed 's/go//')

                        if [ "$INSTALLED_GO" = "$GO_VERSION" ]; then
                            echo "Go $GO_VERSION already installed, skipping."
                        else
                            echo "Installing Go $GO_VERSION (found: ${INSTALLED_GO:-none})..."
                            wget -q "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
                                -O /tmp/go.tar.gz

                            mkdir -p ${WORKSPACE}/tools
                            rm -rf ${WORKSPACE}/tools/go
                            tar -C ${WORKSPACE}/tools -xzf /tmp/go.tar.gz
                            rm /tmp/go.tar.gz

                            export PATH="${WORKSPACE}/tools/go/bin:${PATH}"
                            echo "Go installed: $(go version)"
                        fi

                        # ── Terraform ────────────────────────────────────────
                        INSTALLED_TF=$(terraform version 2>/dev/null | head -1 | awk '{print $2}' | sed 's/v//')

                        if [ "$INSTALLED_TF" = "$TF_VERSION" ]; then
                            echo "Terraform $TF_VERSION already installed, skipping."
                        else
                            echo "Installing Terraform $TF_VERSION (found: ${INSTALLED_TF:-none})..."
                            wget -q \
                                "https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip" \
                                -O /tmp/terraform.zip

                            mkdir -p ${WORKSPACE}/bin
                            unzip -o /tmp/terraform.zip -d ${WORKSPACE}/bin/
                            rm /tmp/terraform.zip

                            echo "Terraform installed: $(${WORKSPACE}/bin/terraform version)"
                        fi

                        # ── go-junit-report ──────────────────────────────────
                        if ! command -v go-junit-report &> /dev/null; then
                            echo "Installing go-junit-report..."
                            go install github.com/jstemmer/go-junit-report/v2@latest
                        else
                            echo "go-junit-report already present."
                        fi
                    '''
                }

                // ── Setup credentials ─────────────────────────────────────────
                stage('Setup credentials') {
                    withCredentials([
                        usernamePassword(
                            credentialsId: 'tf_automation',
                            usernameVariable: 'HOST_A_USER',
                            passwordVariable: 'HOST_A_PASS'
                        ),
                        usernamePassword(
                            credentialsId: 'tf_automation',
                            usernameVariable: 'HOST_B_USER',
                            passwordVariable: 'HOST_B_PASS'
                        )
                    ]) {
                        env.NIOS_USERNAME        = HOST_A_USER
                        env.NIOS_PASSWORD        = HOST_A_PASS
                        env.NIOS_MEMBER_USERNAME = HOST_B_USER
                        env.NIOS_MEMBER_PASSWORD = HOST_B_PASS
                    }
                    echo "Credentials loaded for ${env.NIOS_HOST_URL} and ${env.NIOS_MEMBER_URL}"
                }

                // ── Integration config ────────────────────────────────────────
                stage('Integration config') {
                    sh '''#!/bin/bash
                        set -euo pipefail

                        rm -f pipeline.env integration_test_setup.log
                        go run ./internal/testdata/integration_test_setup.go 2>&1 | tee integration_test_setup.log

                        if grep -E '(^|[[:space:]])Error( |:)' integration_test_setup.log >/dev/null; then
                            echo "integration_test_setup.go reported an error."
                            exit 1
                        fi

                        if [ ! -s pipeline.env ]; then
                            echo "pipeline.env was not generated or is empty."
                            exit 1
                        fi
                    '''

                    // Inject pipeline.env into Jenkins env
                    def envVars = readFile('pipeline.env').trim().split('\n')
                    envVars.each { line ->
                        def parts = line.split('=', 2)
                        if (parts.length == 2) {
                            env.setProperty(parts[0], parts[1])
                        }
                    }
                }

                // ── Test stages ───────────────────────────────────────────────
                sh 'mkdir -p $WORKSPACE/test-results'
                runTestStage('acl',             'internal/service/acl',             '5m')
                runTestStage('cloud',           'internal/service/cloud',           '5m')
                runTestStage('dhcp',            'internal/service/dhcp',            '80m')
                runTestStage('discovery',       'internal/service/discovery',       '10m')
                runTestStage('dns',             'internal/service/dns',             '80m')
                runTestStage('dtc',             'internal/service/dtc',             '30m')
                runTestStage('grid',            'internal/service/grid',            '30m')
                runTestStage('ipam',            'internal/service/ipam',            '80m')
                runTestStage('microsoft',       'internal/service/microsoft',       '30m')
                runTestStage('misc',            'internal/service/misc',            '30m')
                runTestStage('notification',    'internal/service/notification',    '30m')
                runTestStage('parentalcontrol', 'internal/service/parentalcontrol', '30m')
                runTestStage('rir',             'internal/service/rir',             '5m')
                runTestStage('rpz',             'internal/service/rpz',             '45m')
                runTestStage('security',        'internal/service/security',        '30m')
                runTestStage('smartfolder',     'internal/service/smartfolder',     '5m')

            } catch (err) {
                currentBuild.result = 'FAILURE'
                echo "Pipeline failed: ${err}"
                throw err
            } finally {
                // ── Post-build actions (always run) ───────────────────────────
                stage('Cleanup') {
                    echo 'Running cleanup...'
                    //sh 'go run ./internal/testdata/integration_test_cleanup.go || true'
                }

                // Publish JUnit + archive logs
                junit allowEmptyResults: true, testResults: 'test-results/**-junit.xml'
                archiveArtifacts artifacts: 'test-results/*.txt', allowEmptyArchive: true

                if (currentBuild.result == 'FAILURE') {
                    echo """
                    ❌ One or more tests failed. Check the Test Results tab above.
                    Grid Master: ${env.NIOS_HOST_URL}
                    Grid Member: ${env.NIOS_MEMBER_URL}
                    """
                } else {
                    echo """
                    ✅ All tests passed.
                    Grid Master: ${env.NIOS_HOST_URL}
                    Grid Member: ${env.NIOS_MEMBER_URL}
                    """
                }
            }

        } // timeout
    } // timestamps
} // node