
namespace go useraction

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct User {
  1:i64 user_id
  2:string user_name
  3:i64 follow_count
  4:i64 follower_count
  5:optional bool is_follow
}

struct GetUserRequest{
  1:string user_id
}

struct GetUserResponse{
  1:User user
  2:BaseResp base_resp
}


struct Video{
  1:i64 video_id
  2:User user
  3:string play_url
  4:string cover_url
  5:i64 favorite_count
  6:i64 comment_count
  7:bool is_favorite
  8:string title
}

struct PublishListRequest{
  1:i64 user_id
}

struct PublishListResponse{
  1:Video video
  2:BaseResp base_resp
}

struct FavoriteActionRequest{
  1:i64 user_id
  2:i64 video_id
  3:i32 action_type
}

struct FavoriteActionResponse{
  1:BaseResp base_resp
}

struct FavoriteListRequest{
  1:i64 user_id
}

struct FavoriteListResponse{
  1:BaseResp base_resp
  2:list<Video> video_list
}

struct Comment{
  1:i64 comment_id
  2:User user
  3:string content
  4:string create_date
}

struct CommentActionRequest{
  1:i64 user_id
  2:string video_id
  3:string action_type
  4:string comment_text
  5:i64 comment_id
}

struct CommentActionResponse{
  1:BaseResp base_resp
  2:Comment comment
}

struct CommentListRequest{
  1:i64 video_id
}

struct CommentListResponse{
  1:BaseResp base_resp
  2:list<Comment> comment_list
}

struct RelationActionRequest{
  1:string user_id
  2:string to_user_id
  3:i32 action_type
}

struct RelationActionResponse{
  1:BaseResp base_resp
}

struct RelationFollowListRequest{
  1:string user_id
}

struct RelationFollowListResponse{
  1:BaseResp base_resp
  2:list<User> user_list
}

struct RelationFollowerListRequest{
  1:string user_id
}

struct RelationFollowerListResponse{
  1:BaseResp base_resp
  2:list<User> user_list
}

service UserActionService {
      GetUserResponse GetUser(1:GetUserRequest req)
      PublishListResponse PublishList(1:PublishListRequest req)
      FavoriteActionResponse FavoriteAction(1:FavoriteActionRequest req)
      FavoriteListResponse FavoriteList(1:FavoriteListRequest req)
      CommentActionResponse CommentAction(1:CommentActionRequest req)
      CommentListResponse CommentList(1:CommentListRequest req)
      RelationActionResponse RelationAction(1:RelationActionRequest req)
      RelationFollowListResponse RelationFollowList(1:RelationFollowListRequest req)
      RelationFollowerListResponse RelationFollowerList(1:RelationFollowerListRequest req)
}