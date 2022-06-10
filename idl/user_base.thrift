namespace go userbase

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}



struct CreateUserRequest {
    1:string user_name
    2:string password
    3:i64 user_id
}

struct CreateUserResponse {
    1:BaseResp base_resp
}


struct CheckUserRequest{
    1:string user_name
    2:string password
}

struct CheckUserResponse{
    1:i64 user_id
    2:BaseResp base_resp
}


service UserBaseService {
    CreateUserResponse CreateUser(1:CreateUserRequest req)
    CheckUserResponse CheckUser(1:CheckUserRequest req)
}