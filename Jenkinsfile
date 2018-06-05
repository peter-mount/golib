// Build properties
properties([
  buildDiscarder(logRotator(artifactDaysToKeepStr: '', artifactNumToKeepStr: '', daysToKeepStr: '', numToKeepStr: '10')),
  disableConcurrentBuilds(),
  disableResume(),
  pipelineTriggers([
    cron('H H * * *')
  ])
])

stage( 'prepare' ) {
  sh 'docker build -t golib:build --target build'
}

[
  'codec', 'kernel', 'rabbitmq', 'rest', 'statistics'
].each {
  moduleName ->
    stage( moduleName ) {
      sh 'docker build' +
         ' -t golib:' + moduleName +
         ' --build-arg moduleName=' + moduleName +
         ' .'
    }
}