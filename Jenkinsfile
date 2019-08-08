pipeline {
    agent any

    triggers {
        cron('0 12 * * *')
    }
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
