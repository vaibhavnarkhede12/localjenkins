pipeline {
  agent any
    
  tools {go "golang"}
    
  parameters{
    choice(
    name:'language',
      choices:[
        'python',
        'golang',
        'javaa',
        'engg'
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
        echo "displaying testversion stage for brnach ${BRANCH_NAME}  - BUILD URL ${BUILD_URL}"
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
        echo "executing the ${BRANCH_NAME} in BUILD STAGE - BUILD URL ${BUILD_URL}"
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
        echo "executing the ${BRANCH_NAME} in DEPLOY STAGE  - BUILD URL ${BUILD_URL}"
        sh 'go run serve.go'
      }
    }
  }
}
