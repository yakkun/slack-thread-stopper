slack-thread-stopper
===

![利用イメージ](./docs/image.png "利用イメージ")

スレッド利用を非推奨にする Slack Bot です。この Bot がいるチャンネルではスレッドを利用した場合に非推奨であると警告されます。

## Requirement

- Go 1.14+
- Slack

## Setup

Slack のソケットモードを利用しています。

環境変数にて、 Slack の Token が 2 種類、必要です。

- `SLACK_BOT_TOKEN`: `xoxb-` から始まる Bot トークン
- `SLACK_APP_TOKEN`: `xapp-` から始まる App トークン

`.env` ファイルが利用できます。 [.env.sample](/.env.sample) をコピーしてご利用ください。
