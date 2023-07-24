### Core libraries used in rew3 projects

#### Folder Structure
Reference: https://github.com/SKF/go-utility
Reference: https://github.com/eiicon-company/go-core

#### Creating Tags and Publishing Build to Github
Reference: https://gist.github.com/danielestevez/2044589

Current version of build is `v1.0.1`, and branch is `v1.0.1-branch`

1) Create a branch with the tag, currently existing branch is v1.0.1-branch
	`git branch {tagname}-branch` (only create new branch if needed for new version)
	`git checkout {tagname}-branch e.g. v1.0.1-branch`

2) Include your changes or merge latest changes to this tag branch. 
	
3) Delete and recreate the tag locally
	`git tag -d {tagname} e.g. git tag -d v1.0.1`
	`git tag {tagname} e.g. git tag v1.0.1`

4) Delete and recreate the tag remotely
	`git push origin :{tagname} e.g. git push origin :v1.0.1` // deletes original remote tag
	`git push origin {tagname} e.g. git push origin v1.0.1` // creates new remote tag
		
5)  Update local repository with the updated tag.
	git fetch --tags
	
Note, we have issue in version 1.0.0, so we use version 1.0.1 for rew3-base. 

Update on git tag:
- Simply update tag `git tag -f <tag_name_to_update> <hash_code_new_commit>`
- Delete and Push to repository `git push origin {tagname} or git push --tags`

Reference: https://www.toolsqa.com/git/git-delete-tag/#:~:text=Similar%20to%20this%2C%20sometimes%2C%20we,depicting%20that%20the%20stable%20Version1.


### Generating Proto Files:
Install Protoc - `https://google.github.io/proto-lens/installing-protoc.html`  
Download protoc-gen-go  - `go get -u github.com/golang/protobuf/protoc-gen-go`

Usage:  
`protoc --go_out=. --go-grpc_out=. path/to/your/protofile.proto`  
 Replace `path/to/your/protofile.proto` with the path to your `.proto` file.

Setup Environment Variable:  
- Ensure `GOPATH` and `GOBIN` env variable set. If not, set it: `export GOPATH=$HOME/go`, `export GOBIN=$GOPATH/bin`.
- Ensure `GOPATH` exists in system's `Path` variable (`echo $PATH`). If not, export it: `export PATH=$PATH:$(go env GOPATH)/bin`

You might also need this:  
`go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`  
`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc`
