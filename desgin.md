# API 设计

1. Mysql 的数据结构表。 分为两个表设计。 

	- follow 
	- follow_by 
	- friend_request

2. 具体的表结构


follow(id, user_id, follow_by_user_id, created_at)

followby(id, user_id, follower_id) 


friend(id, user_1_id, user_2_id, created_at)


fredn_request(id, from_user_id, to_user_id, status, created_at) 

status = （0， 1，2） ， 正在申请， 已经完成，忽略

3. 说明

	1. 关注系统设计。 follow 和 follow_by 进行解耦。 好处， 坏处。 
	2. 好友请求表设计。 



4. 数据流转流程。 


	1. 申请好友。  A user_id = 1  想申请  B  user_id = 2 的好友请求

		往 freind_request 表写一个数据。 (1, from_user_id = 1, to_user_id = 2, status = 0, created_at = now()) 

	2. 获取申请列表。 B 用户，需要获取哪些人，想加自己的好友。 
		1. select * from friend_request where to_user_id = 2 (b 用户的 id) and status = 0  
		2. status = 0 , 还没处理，或者是正在处理中的一个状态。 

	3. 好友双向删除。 
		一、 方案一， 只用一个 freind() ,  
		二、 方案二， 拆分为两个表， follow , follow_by 。

		1. 好处， x, 坏处 查。 扩展性差。 2 ， interct 交集。 


		只用一个表。 friend
			select * from frined where user_id_1 = 1 or user_id_2 = 1 ; (查处所有 1 的好友） 
			(建立好友关系 需要比较这两个人的 ID user_id_1 > user_id_3) 

		delete from friend where user_id_1 = xx and user_id_2 = xxx 。 （ A, B ） 

	4. 好友列表。 

		 select * from frined where user_id_1 = 1 or user_id_2 = 1 ; (查处所有 1 的好友）



5. 后端的 API  设计。 



添加好友、删除好友

input {
	user_id: xxxx 
	friend_user_id: xxxx, 
	action: (0, 1)
		0: 申请好友
		1: 删除好友
}

output { }


通过好友申请、忽略好友申请。 可以合并为一个接口。 

req {

	user_id : xxx, 
	friend_user_id xxxx, 
	action: (0, 1) 
		0: 通过
		1: 忽略，拒绝
}

resp { empty / friend_user info }



好友列表。 

req { 
	user_id 
}


resp {

	[]freind_user_ids,  
	offset: 50, 
	totalNum: 1000, 
}

pagination, user_details_info 

cache 。 user_id -> 内存的值 -> redis key -> mysql 。 

- rpc 的设计基本完成。 



- 后端的伪代码。 




1、项目是 部署在 k8s pod, 分片部署， 有一个 API  GeteWay 。 


前端 http Post  域名 xxxxx.com/xxxxService.ApplyFriendRequest 

req {

        user_id : xxx,
        friend_user_id xxxx,
        action: (0, 1)
                0: 通过
                1: 忽略，拒绝
}

grpc 。 unary  


lb = nginx ， -> k8s 的一个网关中。  

网关服务， API Gateway 。 熔断、限流、鉴权。 

根据 token ， 鉴权完毕。 


API Getaway 根据方法名称， ApplyFreindRequest  。 找到我们后背的一个微服务 Freind Service 。 

	API Getaway  connection pool ( friend service, xxx service ) （服务注册、服务发现、配置中心 ）进行管理

配置 rounter 表，反向找到 service 的 conn 。 



请求就打到了 friend service 的 pod 。 （负载均衡） 

friend request service 的进程 。 
好友申请。 


fucn (s *FriendServer) ApplyFreind(ctx, req) resp { 
	// req 请求校验工作。 

	switch req.Action {
	case 0: 
		// 添加好友
		err := db.AddFriend(req.UserId, req.FriendId) 
		if err ! {
			return err 
		}
	case 1: 
		// 删除好友
		db.DeleteFriend(req.UserId, req.FreindId)
	}
	return resp{{ 
}

resp {
	
	etc: xxxx  tag kv-
}

