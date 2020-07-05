// Code generated by github.com/Yamashou/gqlgenc, DO NOT EDIT.

package gen

import (
	"context"

	"github.com/Yamashou/gqlgenc/client"
)

type Client struct {
	Client *client.Client
}
type UserFragment struct {
	Name     *string
	Username string
}
type GetUsers struct {
	QueryUser []*struct {
		Name     *string
		Username string
		Tasks    []*struct {
			Title     string
			Completed bool
		}
	}
}
type AddTasksPayload struct {
	AddTask *struct {
		NumUids *int
		Task    []*struct {
			ID    string
			Title string
			User  UserFragment
		}
	}
}
type AddUsersMutationPayload struct {
	AddUser *struct {
		NumUids *int
		User    []*UserFragment
	}
}
type DeleteUserMutationPayload struct {
	DeleteUser *struct {
		NumUids *int
		Msg     *string
	}
}
type DeleteTaskMutationPayload struct {
	DeleteTask *struct {
		NumUids *int
		Msg     *string
	}
}

const GetUsersQuery = `query GetUsers {
	queryUser {
		... UserFragment
		tasks {
			title
			completed
		}
	}
}
fragment UserFragment on User {
	name
	username
}
`

func (c *Client) GetUsers(ctx context.Context, httpRequestOptions ...client.HTTPRequestOption) (*GetUsers, error) {
	vars := map[string]interface{}{}

	var res GetUsers
	if err := c.Client.Post(ctx, GetUsersQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const AddTasksQuery = `mutation AddTasks ($input: [AddTaskInput!]!) {
	addTask(input: $input) {
		numUids
		task {
			id
			title
			user {
				... UserFragment
			}
		}
	}
}
fragment UserFragment on User {
	name
	username
}
`

func (c *Client) AddTasks(ctx context.Context, input []*AddTaskInput, httpRequestOptions ...client.HTTPRequestOption) (*AddTasksPayload, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res AddTasksPayload
	if err := c.Client.Post(ctx, AddTasksQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const AddUsersMutationQuery = `mutation AddUsersMutation ($input: [AddUserInput!]!) {
	addUser(input: $input) {
		numUids
		user {
			... UserFragment
		}
	}
}
fragment UserFragment on User {
	name
	username
}
`

func (c *Client) AddUsersMutation(ctx context.Context, input []*AddUserInput, httpRequestOptions ...client.HTTPRequestOption) (*AddUsersMutationPayload, error) {
	vars := map[string]interface{}{
		"input": input,
	}

	var res AddUsersMutationPayload
	if err := c.Client.Post(ctx, AddUsersMutationQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const DeleteUserMutationQuery = `mutation DeleteUserMutation ($username: String!) {
	deleteUser(filter: {username:{eq:$username}}) {
		numUids
		msg
	}
}
`

func (c *Client) DeleteUserMutation(ctx context.Context, username string, httpRequestOptions ...client.HTTPRequestOption) (*DeleteUserMutationPayload, error) {
	vars := map[string]interface{}{
		"username": username,
	}

	var res DeleteUserMutationPayload
	if err := c.Client.Post(ctx, DeleteUserMutationQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}

const DeleteTaskMutationQuery = `mutation DeleteTaskMutation ($id: ID!) {
	deleteTask(filter: {id:[$id]}) {
		numUids
		msg
	}
}
`

func (c *Client) DeleteTaskMutation(ctx context.Context, id string, httpRequestOptions ...client.HTTPRequestOption) (*DeleteTaskMutationPayload, error) {
	vars := map[string]interface{}{
		"id": id,
	}

	var res DeleteTaskMutationPayload
	if err := c.Client.Post(ctx, DeleteTaskMutationQuery, &res, vars, httpRequestOptions...); err != nil {
		return nil, err
	}

	return &res, nil
}