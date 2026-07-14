#!/usr/bin/env groovy
/**
 * Required environment variables (set by Jenkinsfile before loading):
 *   SIGNAL_GIST_ID   – Gist ID to write signals into
 *   SIGNAL_TOKEN     – GitHub PAT with gist + actions:write scope
 */

/**
 * Push a JSON payload to the signal Gist as an immutable per-stage file.
 * Each call creates a NEW file: ci-{sha}-{build#}-{stage}.json
 */
private void _push(Map payload, Map extraFiles = [:]) {
    def gistId      = params.SIGNAL_GIST_ID?.trim() ?: env.SIGNAL_GIST_ID ?: ''
    def token       = params.SIGNAL_TOKEN?.trim()   ?: env.SIGNAL_TOKEN   ?: ''
    def sha         = params.GIT_COMMIT?.trim()     ?: env.GIT_COMMIT     ?: ''
    def buildNumber = env.BUILD_NUMBER ?: '0'

    if (!gistId || !token) {
        echo "⚠  SIGNAL_GIST_ID or SIGNAL_TOKEN not set — skipping signal push."
        return
    }

    payload.sha          = sha
    payload.build_number = buildNumber
    payload.timestamp    = new Date().format("yyyy-MM-dd'T'HH:mm:ss'+00:00'", TimeZone.getTimeZone('UTC'))
    payload.jenkins_url  = env.BUILD_URL ?: ''

    def stage         = payload.stage ?: 'unknown'
    def stageFileName = "ci-${sha}-${buildNumber}-${stage}.json"

    def files = [:]
    files[stageFileName] = [content: groovy.json.JsonOutput.toJson(payload)]
    extraFiles.each { fname, content ->
        files[fname.toString()] = [content: content]
    }
    def jsonBody = groovy.json.JsonOutput.toJson([files: files])

    writeFile file: '.signal_payload.json', text: jsonBody

    def http = sh(
        script: """
            curl -sS -o /dev/null -w "%{http_code}" \\
                -X PATCH \\
                -H "Authorization: token ${token}" \\
                -H "Accept: application/vnd.github+json" \\
                -H "Content-Type: application/json" \\
                --data @.signal_payload.json \\
                "https://api.github.com/gists/${gistId}"
        """,
        returnStdout: true
    ).trim()

    sh "rm -f .signal_payload.json"

    if (http != '200') {
        echo "⚠  Gist PATCH returned HTTP ${http}."
    } else {
        echo "📡  Signal pushed → ${stageFileName}"
    }
}

/**
 * Trigger the GitHub Actions integration-tests.yml workflow via workflow_dispatch.
 * Jenkins calls this after tests and email have completed (in the finally block).
 *
 * @param repoUser  GitHub org/user (e.g. 'infobloxopen')
 * @param branch    Branch ref to dispatch on (e.g. 'main')
 */
void triggerGitHubAction(String repoUser, String branch) {
    def token       = params.SIGNAL_TOKEN?.trim() ?: env.SIGNAL_TOKEN ?: ''
    def sha         = env.GIT_COMMIT   ?: ''
    def buildNumber = env.BUILD_NUMBER ?: '0'
    def gistId      = params.SIGNAL_GIST_ID?.trim() ?: env.SIGNAL_GIST_ID ?: ''

    if (!token) {
        echo "⚠  SIGNAL_TOKEN not set — cannot trigger GitHub Actions workflow."
        return
    }

    def payload = groovy.json.JsonOutput.toJson([
        ref   : branch,
        inputs: [
            sha         : sha,
            build_number: buildNumber,
            gist_id     : gistId
        ]
    ])

    writeFile file: '.dispatch_payload.json', text: payload

    def http = sh(
        script: """
            curl -sS -o /dev/null -w "%{http_code}" \\
                -X POST \\
                -H "Authorization: token ${token}" \\
                -H "Accept: application/vnd.github+json" \\
                -H "Content-Type: application/json" \\
                --data @.dispatch_payload.json \\
                "https://api.github.com/repos/${repoUser}/terraform-provider-nios/actions/workflows/integration-tests.yml/dispatches"
        """,
        returnStdout: true
    ).trim()

    sh "rm -f .dispatch_payload.json"

    if (http == '204') {
        echo "✅  GitHub Actions workflow triggered for commit ${sha} (build #${buildNumber})."
    } else {
        echo "⚠  GitHub Actions dispatch returned HTTP ${http}. Check SIGNAL_TOKEN has actions:write scope."
    }
}

/**
 * Send a SUCCESS signal for a stage.
 *
 * @param stageName       matches the stage GHA looks for (e.g. 'env-setup', 'tests-complete')
 * @param opts.message    optional human-readable note
 * @param opts.junit_files  map of filename → XML content to upload alongside the signal
 */
void success(String stageName, Map opts = [:]) {
    def sha         = env.GIT_COMMIT   ?: ''
    def buildNumber = env.BUILD_NUMBER ?: '0'
    def extraFiles  = [:]
    (opts.junit_files ?: [:]).each { name, content ->
        extraFiles["ci-junit-${sha}-${buildNumber}-${name}".toString()] = content
    }
    _push([
        status   : 'success',
        stage    : stageName,
        ref      : env.GIT_BRANCH ?: '',
        message  : opts.message ?: '',
        artifacts: opts.artifacts ?: []
    ], extraFiles)
}

/**
 * Send an ERROR signal — GHA reads this and logs the failure.
 *
 * @param stageName   the stage where the failure occurred
 * @param message     human-readable error description
 */
void error(String stageName, String message) {
    _push([
        status   : 'error',
        stage    : stageName,
        ref      : env.GIT_BRANCH ?: '',
        message  : message,
        artifacts: []
    ])
}

/**
 * Send an ABORTED signal.
 */
void aborted(String stageName, String message = 'Pipeline was aborted') {
    _push([
        status   : 'aborted',
        stage    : stageName,
        ref      : env.GIT_BRANCH ?: '',
        message  : message,
        artifacts: []
    ])
}

/**
 * Send a STARTED signal immediately after checkout.
 * Writes ci-{sha}-{build#}-job-started.json — GHA uses this for audit/validation.
 */
void started() {
    _push([
        status   : 'started',
        stage    : 'job-started',
        ref      : env.GIT_BRANCH ?: '',
        message  : "Jenkins started build for commit ${env.GIT_COMMIT ?: 'unknown'}",
        artifacts: []
    ])
}

return this
