pipeline {
    agent any

    // ── Parameters ────────────────────────────────────────────────────────────
    // Shown as form fields when "Build with Parameters" is triggered.
    parameters {
        // TODO : Remove hardcoded defaults
        string(
            name: 'GM_URL',
            defaultValue: 'https://172.28.82.73',
            description: 'URL For the GRID Master'
        )
        string(
            name: 'MEMBER_URL',
            defaultValue: 'https://172.28.82.16',
            description: 'URL for the GRID Member'
        )
        string(
            name: 'GO_VERSION',
            defaultValue: '1.25.1',
            description: 'Go version to install e.g. 1.22.5'
        )
        string(
            name: 'TF_VERSION',
            defaultValue: '1.8.0',
            description: 'Terraform version to install e.g. 1.8.0'
        )
    }

    // ── Pipeline-wide environment variables ───────────────────────────────────
    // String params are promoted to env vars here so every stage and sh step
    // can use $NIOS_HOST_URL / $NIOS_MEMBER_URL without extra plumbing.
    environment {
        NIOS_HOST_URL = "${params.GM_URL}"
        NIOS_MEMBER_URL = "${params.MEMBER_URL}"
        NIOS_WAPI_VERSION = "v2.13.6"
        TF_ACC     = "1"
        GO_VERSION  = "${params.GO_VERSION}"
        TF_VERSION  = "${params.TF_VERSION}"
        PATH        = "/usr/local/go/bin:${HOME}/go/bin:${env.PATH}"

        // withCredentials() below injects NIOS_USERNAME, NIOS_PASSWORD, etc.
        // Declaring them here as empty strings lets the IDE / linter know they exist.
        HOST_A_USER = ''
        HOST_A_PASS = ''
        HOST_B_USER = ''
        HOST_B_PASS = ''
    }

    options {
        // Show timestamps next to every log line — very useful for long test runs.
        timestamps()
        // Fail the whole build if it takes longer than 120 minutes.
        timeout(time: 120, unit: 'MINUTES')
    }

    stages {

        // Checkout Step to add the code to the Jenkins workspace 
        stage('Checkout') {
            // steps {
            //     sh 'cp -r /Users/unasra/go/src/github.com/infobloxopen/terraform-provider-nios/. ${WORKSPACE}/ \
            //     --exclude=.git'
            //     sh '''
            //         cp -r /Users/unasra/go/src/github.com/infobloxopen/terraform-provider-nios/. ${WORKSPACE}/ \
            //             --exclude='.git'
            //     '''
            // }
//             steps {
//                 git branch: 'fix_tests',
//                     url: 'https://github.com/unasra/terraform-provider-nios.git',
//             }
            steps {
                checkout([
                    $class: 'GitSCM',
                    branches: [[name: '*/fix_tests']],
                    userRemoteConfigs: [[
                        url: 'file:///Users/unasra/go/src/github.com/infobloxopen/terraform-provider-nios'
                    ]],
                    extensions: [[$class: 'CleanBeforeCheckout']]
                ])
            }
        }

        // Setup the ToolChain
        stage('Install toolchain') {
            steps {
                sh '''
                    # ── Go ───────────────────────────────────────────────────────
                    INSTALLED_GO=$(go version 2>/dev/null | awk '{print $3}' | sed 's/go//')

                    if [ "$INSTALLED_GO" = "$GO_VERSION" ]; then
                        echo "Go $GO_VERSION already installed, skipping."
                    else
                        echo "Installing Go $GO_VERSION (found: ${INSTALLED_GO:-none})..."
                        # wget -q "https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
                        #     -O /tmp/go.tar.gz
                        # rm -rf /usr/local/go
                        # tar -C /usr/local -xzf /tmp/go.tar.gz
                        # rm /tmp/go.tar.gz
                        echo "Go installed: $(go version)"
                    fi

                    # ── Terraform ────────────────────────────────────────────────
                    INSTALLED_TF=$(terraform version 2>/dev/null | head -1 | awk '{print $2}' | sed 's/v//')

                    if [ "$INSTALLED_TF" = "$TF_VERSION" ]; then
                        echo "Terraform $TF_VERSION already installed, skipping."
                    else
                        echo "Installing Terraform $TF_VERSION (found: ${INSTALLED_TF:-none})..."
                        wget -q \
                            "https://releases.hashicorp.com/terraform/${TF_VERSION}/terraform_${TF_VERSION}_linux_amd64.zip" \
                            -O /tmp/terraform.zip

                        # Install into workspace/bin — Jenkins always has write access here
                        mkdir -p ${WORKSPACE}/bin
                        unzip -o /tmp/terraform.zip -d ${WORKSPACE}/bin/
                        rm /tmp/terraform.zip

                        echo "Terraform installed: $(${WORKSPACE}/bin/terraform version)"
                    fi

                    # ── go-junit-report ──────────────────────────────────────────
                    if ! command -v go-junit-report &> /dev/null; then
                        echo "Installing go-junit-report..."
                        go install github.com/jstemmer/go-junit-report/v2@latest
                    else
                        echo "go-junit-report already present."
                    fi
                '''
            }
        }

        // ── Stage 1: Inject credentials ───────────────────────────────────────
        // This stage exists only to pull secrets out of the Jenkins store and
        // write them into env vars that survive into later stages.
        // NOTE: Jenkins masks these values in the console log automatically.
        stage('Setup credentials') {
            steps {
                withCredentials([
                    usernamePassword(
                        credentialsId: 'host-a-creds',
                        usernameVariable: 'HOST_A_USER',
                        passwordVariable: 'HOST_A_PASS'
                    ),
                    usernamePassword(
                        credentialsId: 'host-b-creds',
                        usernameVariable: 'HOST_B_USER',
                        passwordVariable: 'HOST_B_PASS'
                    )
                ]) {
                    // Export them as real env vars so subsequent stages pick them up.
                    // (withCredentials scope ends here, but script{} exports persist
                    //  within the same node/agent for the build lifetime.)
                    script {
                        env.NIOS_USERNAME = HOST_A_USER
                        env.NIOS_PASSWORD = HOST_A_PASS
                        env.NIOS_MEMBER_USERNAME = HOST_B_USER
                        env.NIOS_MEMBER_PASSWORD = HOST_B_PASS
                    }
                    echo "Credentials loaded for ${NIOS_HOST_URL} and ${NIOS_MEMBER_URL}"
                }
            }
        }

        // ── Stage 2: Integration setup ────────────────────────────────────────
        // Runs config_file.go. If ANY function in that file returns a non-zero
        // exit code, this stage fails and the pipeline stops — tests don't run
        // against a broken environment.
        stage('Integration config') {
            steps {
                sh '''
                    set -euo pipefail

                    # Run Setup Script
                    rm -f pipeline.env integration_test_setup.log
                    go run ./internal/testdata/integration_test_setup.go 2>&1 | tee integration_test_setup.log

                    # Fail the stage if setup script printed an error but exited 0
                    if grep -E '(^|[[:space:]])Error( |:)' integration_test_setup.log >/dev/null; then
                        echo "integration_test_setup.go reported an error."
                        exit 1
                    fi

                    # Ensure setup generated environment variables file
                    if [ ! -s pipeline.env ]; then
                        echo "pipeline.env was not generated or is empty."
                        exit 1
                    fi
                '''
                // Read each line from the file and inject into Jenkins env
                script {
                    def envVars = readFile('pipeline.env').trim().split('\n')
                    envVars.each { line ->
                        def (key, value) = line.tokenize('=')
                        env[key] = value
                    }
                }
            }
        }

        // ── Stage 3: Test directories ─────────────────────────────────────────
        // Each subdirectory becomes its own named stage so Jenkins' stage view
        // shows individual pass/fail per directory.
        //
        // Why sequential and not parallel{} blocks?
        // Your tests hit a live host and parallel runs cause connection timeouts.
        // Sequential = one at a time, but failFast:false means a failure in
        // dir_1 does NOT abort dir_2 or dir_3 — all dirs always run.
        //
        // JUnit XML is written per-directory so the final report is granular.
        stage('Run tests') {
            // failFast false → a failed directory does not abort the remaining ones.
            failFast false

            // Replace the list below with your actual test directory names.
            // Each entry becomes a separate visible stage in Jenkins Blue Ocean / Stage View.
            stages {

                stage('Tests: acl') {
                    steps {
                        catchError(buildResult: 'FAILURE', stageResult: 'FAILURE') {
                            sh '''
                                mkdir -p $WORKSPACE/test-results
                                cd internal/service/acl
                                go test -v -count=1 -timeout 5m ./... \
                                    2>&1 | tee $WORKSPACE/test-results/acl.txt
                            '''
                        }
                    }
                    // Convert the raw output to JUnit XML so Jenkins can graph it.
                    post {
                        always {
                            sh 'go-junit-report < $WORKSPACE/test-results/acl.txt > $WORKSPACE/test-results/acl-junit.xml || true'
                        }
                    }
                }

                stage('Tests: cloud') {
                    steps {
                        catchError(buildResult: 'FAILURE', stageResult: 'FAILURE') {
                            sh '''
                                cd internal/service/cloud
                                go test -v -count=1 -timeout 5m ./... \
                                    2>&1 | tee $WORKSPACE/test-results/cloud.txt
                            '''
                        }
                    }
                    post {
                        always {
                            sh 'go-junit-report < $WORKSPACE/test-results/cloud.txt > $WORKSPACE/test-results/cloud-junit.xml || true'
                        }
                    }
                }

                // stage('Tests: dhcp') {
                //     steps {
                //         sh '''
                //             cd internal/service/dhcp
                //             go test -v -count=1 -timeout 30m ./... \
                //                 2>&1 | tee ../../test-results/dhcp.txt
                //         '''
                //     }
                //     post {
                //         always {
                //             sh 'go-junit-report < ../../test-results/dhcp.txt > ../../test-results/dhcp-junit.xml || true'
                //         }
                //     }
                // }

               stage('Tests: discovery') {
                    steps {
                        catchError(buildResult: 'FAILURE', stageResult: 'FAILURE') {
                            sh '''
                                cd internal/service/discovery
                                go test -v -count=1 -timeout 30m ./... \
                                    2>&1 | tee $WORKSPACE/test-results/discovery.txt
                            '''
                        }
                    }
                    post {
                        always {
                            sh 'go-junit-report < $WORKSPACE/test-results/discovery.txt > $WORKSPACE/test-results/discovery-junit.xml || true'
                        }
                    }
                }

                // stage('Tests: dns') {
                //     steps {
                //         sh '''
                //             cd internal/service/dns
                //             go test -v -count=1 -timeout 30m ./... \
                //                 2>&1 | tee ../../test-results/dns.txt
                //         '''
                //     }
                //     post {
                //         always {
                //             sh 'go-junit-report < ../../test-results/dns.txt > ../../test-results/dns-junit.xml || true'
                //         }
                //     }
                // }

//                 stage('Tests: dtc') {
//                     steps {
//                         sh '''
//                             cd internal/service/dtc
//                             go test -v -count=1 -timeout 30m ./... \
//                                 2>&1 | tee ../../test-results/dtc.txt
//                         '''
//                     }
//                     post {
//                         always {
//                             sh 'go-junit-report < ../../test-results/dtc.txt > ../../test-results/dtc-junit.xml || true'
//                         }
//                     }
//                 }

                // stage('Tests: grid') {
                //     steps {
                //         sh '''
                //             cd internal/service/grid
                //             go test -v -count=1 -timeout 30m ./... \
                //                 2>&1 | tee ../../test-results/grid.txt
                //         '''
                //     }
                //     post {
                //         always {
                //             sh 'go-junit-report < ../../test-results/grid.txt > ../../test-results/grid-junit.xml || true'
                //         }
                //     }
                // }

                // stage('Tests: ipam') {
                //     steps {
                //         sh '''
                //             cd internal/service/ipam
                //             go test -v -count=1 -timeout 30m ./... \
                //                 2>&1 | tee ../../test-results/ipam.txt
                //         '''
                //     }
                //     post {
                //         always {
                //             sh 'go-junit-report < ../../test-results/ipam.txt > ../../test-results/ipam-junit.xml || true'
                //         }
                //     }
                // }

                stage('Tests: rir') {
                    steps {
                        catchError(buildResult: 'FAILURE', stageResult: 'FAILURE') {
                            sh '''
                                cd internal/service/rir
                                go test -v -count=1 -timeout 5m ./... \
                                    2>&1 | tee $WORKSPACE/test-results/rir.txt
                            '''
                        }
                    }
                    post {
                        always {
                            sh 'go-junit-report < $WORKSPACE/test-results/rir.txt > $WORKSPACE/test-results/rir-junit.xml || true'
                        }
                    }
                }
                
            }
        }

        // ── Stage 4: Cleanup ──────────────────────────────────────────────────
        // Runs unconditionally — even if tests failed — because it lives in
        // post { always {} } below. Declared as a stage here so you can see it
        // in the Stage View. The actual invocation is in post{} to guarantee it
        // runs after every possible exit path.
        stage('Cleanup') {
            steps {
                echo 'Cleanup will run in post{always} to guarantee execution'
            }
        }
    }

    // ── Post-build actions ────────────────────────────────────────────────────
    post {

        // always{} runs regardless of pass/fail/abort.
        always {
            echo 'Running cleanup...'
            //sh 'go run ./internal/testdata/integration_test_cleanup.go || true'

            // Publish all JUnit XML files. Jenkins aggregates them into the
            // "Test Result" trend graph and shows per-test pass/fail history.
            junit allowEmptyResults: true, testResults: 'test-results/**-junit.xml'

            // Archive the raw text logs as build artifacts (downloadable).
            archiveArtifacts artifacts: 'test-results/*.txt', allowEmptyArchive: true
        }

        success {
            echo """
            ✅ All tests passed.
            Grid Master: ${env.NIOS_HOST_URL}
            Grid Member: ${env.NIOS_MEMBER_URL}
            """
        }

        failure {
            echo """
            ❌ One or more tests failed. Check the Test Results tab above.
            Grid Master: ${env.NIOS_HOST_URL}
            Grid Member: ${env.NIOS_MEMBER_URL}
            """
        }

        unstable {
            // Jenkins marks a build UNSTABLE (yellow) when tests ran but some failed.
            // This is separate from a build ERROR (red), which means the pipeline itself crashed.
            echo '⚠️  Build is unstable — test failures detected. See Test Results tab.'
        }
    }
}