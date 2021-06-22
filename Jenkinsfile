pipeline {
  agent any
    
  tools {go "golang"}
    
  parameters{
    choice(
    name:'language',
      choices:[
        'python',
        'golang',
        'java'
      ],
      description:'select the deployment language')
    choice(
      name:'env',
      choices:[
        'dev',
        'prod'
      ],
     description:'chose your environment'
    )
  }
  stages {
        
   
    stage('testversion') {
      when {
        expression {
          BRANCH_NAME != 'master'
        }
      }
      steps {
        sh 'go version'
        echo "displaying testversion stage for brnach ${BRANCH_NAME}"
//         script{
//           setGitHubPullRequestStatus.message("message from jenkins")
//         }
      }
    }  
      
    stage('build') {
       when {
        expression {
          BRANCH_NAME == 'master'
        }
      }
      steps {
        echo "executing the ${BRANCH_NAME} in BUILD STAGE"
        sh 'go build serve.go'
      }
    }
    
    stage('deploy') {
       when {
        expression {
          BRANCH_NAME == 'master'
        }
      }
      steps {
        echo "executing the ${BRANCH_NAME} in DEPLOY STAGE"
        sh 'go run serve.go'
      }
    }
  }
}
