namespace go video

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}



struct User {
  1:i64 id
  2:string name
  3:i64 follow_count
  4:i64 follower_count
  5:optional bool is_follow
}

struct Video{
  1:i64 id
  2:User author
  3:string play_url
  4:string cover_url
  5:i64 favorite_count
  6:i64 comment_count
  7:bool is_favorite
  8:string title
}


struct GetFeedRequest {
  1:i64 latest_time = 1;
  2:i64 user_id = 2;
}

struct GetFeedResponse {
  1:BaseResp base_resp
  2:list<Video> video_list
  3:i64 next_time
}

struct PublishActionRequest {
  1:i64 user_id
  2:binary data
  3:string title
  4:i64 video_id
}

struct PublishActionResponse {
  1:BaseResp base_resp
}

service VideoService {
  GetFeedResponse GetFeed(1:GetFeedRequest req)
  PublishActionResponse PublishAction(1:PublishActionRequest req)
}