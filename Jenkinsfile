node('docker') {
 
    stage 'Checkout'
        checkout scm
    stage 'Build & UnitTest'
        sh "docker build -t sofikul/test-task-control:B1 -f Dockerfile ."
}