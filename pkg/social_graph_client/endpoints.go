package social_graph_client

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type SimpleResponse struct {
	Err string
}

func decodeSimpleResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp SimpleResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeGetFollowingResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp getFollowingResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

func decodeGetFollowersResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp getFollowersResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}

type EndpointSet struct {
	FollowEndpoint       endpoint.Endpoint
	UnfollowEndpoint     endpoint.Endpoint
	GetFollowingEndpoint endpoint.Endpoint
	GetFollowersEndpoint endpoint.Endpoint
}

type FollowRequest struct {
	Followed string
	Follower string
}

type UnfollowRequest struct {
	Followed string
	Follower string
}

type getByUserNameRequest struct {
	Username string
}

type getFollowingResponse struct {
	Following map[string]bool
	Err       string
}

type getFollowersResponse struct {
	Followers map[string]bool
	Err       string
}

func (s EndpointSet) Follow(followed string, follower string) (err error) {
	resp, err := s.FollowEndpoint(context.Background(), FollowRequest{Followed: followed, Follower: follower})
	if err != nil {
		return err
	}
	response := resp.(SimpleResponse)

	if response.Err != "" {
		err = errors.New(response.Err)
	}
	return
}

func (s EndpointSet) Unfollow(followed string, follower string) (err error) {
	resp, err := s.UnfollowEndpoint(context.Background(), UnfollowRequest{Followed: followed, Follower: follower})
	if err != nil {
		return err
	}
	response := resp.(SimpleResponse)

	if response.Err != "" {
		err = errors.New(response.Err)
	}
	return
}

func (s EndpointSet) GetFollowing(username string) (following map[string]bool, err error) {
	resp, err := s.GetFollowingEndpoint(context.Background(), getByUserNameRequest{Username: username})
	if err != nil {
		return
	}
	response := resp.(getFollowingResponse)

	if response.Err != "" {
		err = errors.New(response.Err)
	}
	following = response.Following
	return
}

func (s EndpointSet) GetFollowers(username string) (following map[string]bool, err error) {
	resp, err := s.GetFollowersEndpoint(context.Background(), getByUserNameRequest{Username: username})
	if err != nil {
		return
	}
	response := resp.(getFollowersResponse)

	if response.Err != "" {
		err = errors.New(response.Err)
	}
	following = response.Followers
	return
}
