package responses

import "time"

// FailedResponseは標準的なエラー応答をフォーマットします。
// 操作が失敗した場合にいつでも使用されます。
func FailedResponse(msg, description string) map[string]interface{} {
	return map[string]interface{}{
		"error_message": msg,                                                                   // エラーメッセージ
		"result":        "failure",                                                             // 結果は「failure」
		"value":         "",                                                                    // 値は失敗時には空
		"description":   description,                                                           // 説明
		"executed_at":   time.Now().UTC().Add(time.Hour * 9).Format("2006/01/02 15:04:05.000"), // 実行日時（日本標準時）
	}
}

// SuccessWithDataResponseは、データを返す成功したリクエストの応答をフォーマットします。
// アイテムのリストや単一のオブジェクトを返す必要がある場合に使用します。
func SuccessWithDataResponse(data interface{}, msg string) map[string]interface{} {
	return map[string]interface{}{
		"error_message": "",                                                                    // エラーメッセージは成功時には空
		"result":        "success",                                                             // 結果は「success」
		"value":         data,                                                                  // 実際のデータがここに渡されます
		"description":   msg,                                                                   // 説明
		"executed_at":   time.Now().UTC().Add(time.Hour * 9).Format("2006/01/02 15:04:05.000"), // 実行日時（日本標準時）
	}
}

// SuccessResponseは、データを返す必要のない成功したリクエストの応答をフォーマットします。
// 例えば、成功した削除操作や更新操作の後など。
func SuccessResponse(msg string) map[string]interface{} {
	return map[string]interface{}{
		"error_message": "",                                                                    // エラーメッセージは成功時には空
		"result":        "success",                                                             // 結果は「success」
		"value":         "",                                                                    // 値はデータが返されないため空
		"description":   msg,                                                                   // 説明
		"executed_at":   time.Now().UTC().Add(time.Hour * 9).Format("2006/01/02 15:04:05.000"), // 実行日時（日本標準時）
	}
}
