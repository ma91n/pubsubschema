# pubsubschema

Pub/Subメッセージスキーマ機能のGoサンプル実装。

* https://cloud.google.com/pubsub/docs/schemas

## 事前準備

```bash
# スキーマ作成
gcloud beta pubsub schemas create avroschema1 ^
--type=AVRO ^
--definition="{\"type\":\"record\",\"name\":\"Avro\",\"fields\":[{\"name\":\"StringField\",\"type\":\"string\"},{\"name\":\"FloatField\",\"type\":\"float\"},{\"name\":\"BooleanField\",\"type\":\"boolean\"}]}"


# スキーマ表示
gcloud beta pubsub schemas describe avroschema1


# スキーマ検証
gcloud beta pubsub schemas validate-message ^
--message-encoding=json ^
--message={\"StringField\":\"hello\",\"FloatField\":123.45,\"BooleanField\":true} ^
--schema-name=avroschema1

gcloud beta pubsub schemas validate-message ^
--message-encoding=json ^
--message={\"NGField\":\"dummy\"} ^
--schema-name=avroschema1

gcloud beta pubsub schemas validate-message ^
--message-encoding=json ^
--message={\"StringField\":\"hello\",\"FloatField\":123.45,\"BooleanField\":"xxx"} ^
--schema-name=avroschema1

gcloud beta pubsub schemas validate-message ^
--message-encoding=json ^
--message={\"StringField\":\"hello\",\"FloatField\":123.45,\"BooleanField\":true,\"extra\":\"aaa\"} ^
--schema-name=avroschema1


# トピック作成（JSON）
gcloud beta pubsub topics create avrotopic ^
--message-encoding=JSON ^
--schema=avroschema1

# トピック作成（バイナリ）
gcloud beta pubsub topics create avrotopic2 ^
--message-encoding=BINARY ^
--schema=avroschema1


# 環境変数設定
set GOOGLE_APPLICATION_CREDENTIALS=%USERPROFILE%/.gcp/<project-id>-1234567890abc.json
```
