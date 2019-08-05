pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                    chmod +x build.sh
                    ./build.sh
                '''
            }
        }

        stage('Run') {
            steps {
                sh '''
                docker-compose up
                '''
            }
        }
    }
}