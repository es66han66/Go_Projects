package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// func redisKeyMaker(n int) string {
// 	return "streamKey:" + strconv.Itoa(n)
// }

// func deleteRedisKeys(ctx context.Context, redisClient *redis.Client, numberofStreamsToBeDeleted int) {
// 	for i := 1; i <= numberofStreamsToBeDeleted; i++ {
// 		streamKeyName := redisKeyMaker(i)
// 		fmt.Println(streamKeyName)
// 		response := redisClient.Del(ctx, streamKeyName)
// 		fmt.Println("Response from Del is ", response)
// 	}
// }

// func addStreamAndStreamKeyToSet(ctx context.Context, redisClient *redis.Client, numberofStreamsRequired int) {
// 	for i := 1; i <= numberofStreamsRequired; i++ {
// 		streamKeyName := redisKeyMaker(i)
// 		addStream(ctx, redisClient, streamKeyName)
// 		addStreamkeyToSet(ctx, redisClient, streamKeyName)
// 	}
// }

// func addStreamkeyToSet(ctx context.Context, redisClient *redis.Client, key string) {
// 	response1 := redisClient.SAdd(
// 		ctx,
// 		"streamKeySet",
// 		[]string{
// 			key,
// 		},
// 	)
// 	fmt.Println("Response from Sadd is ", response1)
// }

func traverseSetAndApplyCallback(
	ctx context.Context,
	redisClient *redis.Client,
	scanCount int64,
	cb func(ctx1 context.Context, redisClient *redis.Client, streamKeys []string),
) {
	var cursorStart uint64 = 0
	keysReturned, nextCursor := redisClient.SScan(ctx, "streamKeySet", cursorStart, "", scanCount).Val()
	cb(ctx, redisClient, keysReturned)
	for {
		if nextCursor != 0 {
			keysReturned, nextCursor = redisClient.SScan(ctx, "streamKeySet", nextCursor, "", 100).Val()
			cb(ctx, redisClient, keysReturned)
		} else {
			break
		}
	}
}

func getConsumerGroupsForStreamKey(ctx context.Context, redisClient *redis.Client, streamKeys []string) {
	for _, streamKey := range streamKeys {
		consumerGroupsInfo, _ := redisClient.XInfoGroups(ctx, streamKey).Result()
		/*
			returns slice of consumer groups with each element having below format

			Name            string
			Consumers       int64
			Pending         int64
			LastDeliveredID string
			EntriesRead     int64
			Lag             int64
		*/
		for _, consumerGroupInfo := range consumerGroupsInfo {
			fmt.Printf("Consumer group belonging to %s with name %s has lag %d \n", streamKey, consumerGroupInfo.Name, consumerGroupInfo.Lag)
		}
	}
}

// func addStream(ctx context.Context, redisClient *redis.Client, key string) {
// 	response := redisClient.XAdd(ctx, &redis.XAddArgs{
// 		Stream:     key,
// 		NoMkStream: false,
// 		Values: []string{
// 			"name",
// 			"eshan",
// 		},
// 	})
// 	fmt.Println("Response from Xadd is ", response)
// }

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()

	// deleteRedisKeys(ctx, redisClient, 10)

	// addStreamAndStreamKeyToSet(ctx, redisClient, 10)

	// traverseSetusingMembersCommand(ctx, redisClient)

	traverseSetAndApplyCallback(ctx, redisClient, 10, getConsumerGroupsForStreamKey)

}
