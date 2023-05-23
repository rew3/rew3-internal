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
