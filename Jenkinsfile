pipeline {
  agent any
  stages {
    stage('error') {
      steps {
        sh '''#!/bin/bash

source ~/.bash_profile
source /etc/profile

go build main.go'''
      }
    }

  }
}