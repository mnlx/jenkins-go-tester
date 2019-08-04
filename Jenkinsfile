pipeline {
    agent {
        docker { image 'golang:1.11-alpine' }
    }
    stages {
        stage('Test') {
            steps {
                sh 'node version'
            }
        }
    }
}