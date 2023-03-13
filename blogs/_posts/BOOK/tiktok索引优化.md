---
title: tiktok索引优化
mathjax: true
cover: http://81.68.91.70/api/cover
date: 2023-03-12 14:17:43
tags:
categories:
---

## 索引优化

除开主键索引外,共建立了如下的索引

*   users     			username : 唯一索引
*   videos   			created_at : 普通索引
*   likes                   user_id video_id : 联合索引
*   fellows              user_id  follower_id : 联合索引
*   messages         from_user_id to_user_id create_time : 联合索引