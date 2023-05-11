package model

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
	Id          string
	FirstName   string
	LastName    string
	Title       string
	Userpic     string
	RoleName    string
	JobTitle    string
	CompanyName string
}
