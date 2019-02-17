pipeline {
  agent { dockerfile true }
    
  stages {    
   stage('Cloning Git') {
      steps {
        git 'https://github.com/Sofiukl/task-control.git'
      }
    }

    stage('Build Image') {
        steps {
            app = docker.build('sofiukl/test-task-control')
        }
    }

    stage('Test image') {
        steps {
            app.inside {
                sh 'echo "Test passed"'
            }
        }
    }

  }

}