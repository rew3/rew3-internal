package core

/**
 * A class to represent an ''Mini User''
 *
 * @param _id unique identifier of the User
 * @param firstName first name of the User
 * @param lastName last name of the User
 *
 * @author rew3 on 2023/05/10
 * @version 1.0.0
 *
 */

type MiniUser struct {
	_id         string
	firstName   string
	lastName    string
	title       string
	userpic     string
	roleName    string
	jobTitle    string
	companyName string
}
