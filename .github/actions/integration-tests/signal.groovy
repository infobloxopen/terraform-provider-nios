#!/usr/bin/env groovy
/**
 * Required build parameters (passed by the GitHub Action):
 *   SIGNAL_GIST_ID   – Gist ID to write signals into
 *   SIGNAL_TOKEN     – GitHub PAT with gist scope
 */

/**
 * Push a JSON payload to the signal Gist.
 */
private void _push(Map payload, Map extraFiles = [:]) {
    def gistId  = params.SIGNAL_GIST_ID?.trim() ?: env.SIGNAL_GIST_ID ?: ''
    def token   = params.SIGNAL_TOKEN?.trim() ?: env.SIGNAL_TOKEN ?: ''
    def runId   = params.SIGNAL_RUN_ID?.trim() ?: env.SIGNAL_RUN_ID ?: ''
    def sha     = params.GIT_COMMIT?.trim() ?: env.GIT_COMMIT ?: ''

    if (!gistId || !token) {
        echo "⚠  SIGNAL_GIST_ID or SIGNAL_TOKEN not set — skipping signal push."
        return
    }

    payload.run_id = runId
    payload.sha    = sha

    // Preserve created_at written by the GitHub Action initialisation step.
    // Each Jenkins PATCH replaces the file content entirely, so without this
    // the field disappears after the first signal and age-based Gist pruning
    // silently stops working.
    def signalFileName = "ci-signal-${sha}.json"
    try {
        def getOutput = sh(
            script: """
                curl -sS -f \\
                    -H "Authorization: token ${token}" \\
                    -H "Accept: application/vnd.github+json" \\
                    "https://api.github.com/gists/${gistId}"
            """,
            returnStdout: true
        ).trim()
        def existingContent = new groovy.json.JsonSlurper()
            .parseText(getOutput)?.files?.get(signalFileName)?.content
        if (existingContent) {
            def existingCreatedAt = new groovy.json.JsonSlurper()
                .parseText(existingContent)?.created_at
            if (existingCreatedAt) {
                payload.created_at = existingCreatedAt
            }
        }
    } catch (Exception ignored) {
        // Non-fatal — fall through to the default below
    }
    // Fallback: use current UTC time so created_at is always present
    if (!payload.containsKey('created_at')) {
        payload.created_at = new Date().format("yyyy-MM-dd'T'HH:mm:ss'+00:00'", TimeZone.getTimeZone('UTC'))
    }

    def files = [:]
    files[signalFileName] = [content: groovy.json.JsonOutput.toJson(payload)]
    extraFiles.each { fname, content ->
        files[fname.toString()] = [content: content]
    }
    def jsonBody = groovy.json.JsonOutput.toJson([files: files])

    // Write to a temp file to avoid shell quoting issues with nested JSON
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
        echo "⚠  Gist PATCH returned HTTP ${http}. The GitHub Action may time out."
    } else {
        echo "📡  Signal pushed → stage='${payload.stage}' status='${payload.status}'"
    }
}

/**
 * Send a SUCCESS signal for a stage.
 *
 * @param stageName       matches the STAGE_NAME the Action is waiting for
 * @param opts.message    optional human-readable note
 * @param opts.junit_files  map of filename → XML content to upload to the Gist
 */
void success(String stageName, Map opts = [:]) {
    def sha = env.GIT_COMMIT ?: ''
    def extraFiles = [:]
    (opts.junit_files ?: [:]).each { name, content ->
        extraFiles["ci-junit-${sha}-${name}".toString()] = content
    }
    _push([
        status   : 'success',
        stage    : stageName,
        message  : opts.message ?: '',
        artifacts: opts.artifacts ?: []
    ], extraFiles)
}

/**
 * Send an ERROR signal — the GitHub Action will immediately fail with the message.
 *
 * @param stageName   the stage where the failure occurred
 * @param message     human-readable error description
 */
void error(String stageName, String message) {
    _push([
        status   : 'error',
        stage    : stageName,
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
        message  : message,
        artifacts: []
    ])
}

/**
 * Send a STARTED signal — call immediately after checkout so the
 * GitHub Action knows Jenkins has picked up this commit.
 */
void started() {
    _push([
        status   : 'started',
        stage    : 'job-started',
        message  : "Jenkins started build for commit ${env.GIT_COMMIT ?: 'unknown'}",
        artifacts: []
    ])
}

return this
