pipeline {
  agent any
    
  tools {go "golang"}
    
  stages {
        
   
    stage('testversion') {
      steps {
        sh 'go version'
      }
    }  
      
    stage('build') {
      steps {
        sh 'go build serve.go'
      }
    }
  }
}
