Core libraries used in rew3 projects


### Creating Tags and Publishing Build to Github
Reference: https://gist.github.com/danielestevez/2044589

1) Create a branch with the tag, currently existing branch is v1.0.0-branch
	`git branch {tagname}-branch` (only create new branch if needed for new version)
	`git checkout {tagname}-branch e.g. v1.0.0-branch`

2) Include your changes or merge latest changes to this tag branch. 
	
3) Delete and recreate the tag locally
	`git tag -d {tagname} e.g. git tag -d v1.0.0`
	`git tag {tagname} e.g. git tag v1.0.0`

4) Delete and recreate the tag remotely
	`git push origin :{tagname} e.g. git push origin :v1.0.0` // deletes original remote tag
	`git push origin {tagname} e.g. git pull origin :v1.0.0` // creates new remote tag
		
5)  Update local repository with the updated tag.
	git fetch --tags