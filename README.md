# complete-notifier

- Discord の Webhook を使ってタスク完了の通知を送ります
- Send notifications of task completion using Discord's Webhook

## URL settings

### .env

```
DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/xxxxxx/xxxxxx
```

### flag option

```
./complete-notifier -url=https://discord.com/api/webhooks/xxxxxx/xxxxxx "this is message."
```
