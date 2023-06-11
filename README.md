<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px src="https://i.imgur.com/FxL5qM0.jpg" alt="Bot logo"></a>
</p>

<h3 align="center">Openshift Telegram Bot</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center" style="font-size: 18px;">  🤖 Deploy your applications on Openshift without any interaction with the Openshift Web Console/CLI ✨
    <br>
📮 Send your <b>Container Image</b> and get back the 🚀 <b>Live Application URL</b>! 📡
</p>

<br>

## 💭 How it works <a name = "working"></a>

First, it asks the user to enter the **Cluster URL**. The bot then checks if the URL is valid.  
Then asks the user to enter the **API Token**. To fetch the **API Token**, the bot generates a link which redirects the user to the **Display Token** Page of the Openshift Cluster.  
The user then has to login with their credentials, copy the token and send it to the bot.

**The bot then uses the token to authenticate itself with the Openshift Cluster.**

Once authenticated, the bot asks the user to send the **Container Image** and the **PORT** of the application to be deployed.

In the background, the bot then creates a new deployment, service and route in the Openshift Cluster and returns the **Live Application URL** to the user!

The bot uses the [telegram-bot-api](https://go-telegram-bot-api.dev/) package to interact with the Telegram Bot API and the [kubernetes/client-go]("https://github.com/kubernetes/client-go") package to interact with the Openshift Cluster.

**The entire bot is written in Golang 1.20** 🚀

<br>

## 🎥 Demo  <a name = "demo"></a>

Will Upload Soon!!!
<!-- ![Working](<url>) -->

<br>

## 🎈 Usage <a name = "usage"></a>

To start the Bot:

```
/start
```

***Follow the Instructions to authenticate the bot with your Openshift Cluster.***

To deploy a container image on Openshift:

```
/image <container-image-name> | <PORT>
```

To change the default namespace and user:  
*(Default: namespace = "default", user = "kube:admin")*

```
/default <namespace> | <user>
```

<br>

## 🏁 Getting Started <a name = "getting_started"></a>

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

```
- You need to have a Telegram Account. 
- Register a new bot with BotFather. 
- Pass the TELEGRAM_BOT_TOKEN received in a .env file in the root directory.

- You would need a Openshift Cluster to test this bot. 
(You can get a free cluster from Developer Sandbox for Red Hat OpenShift!)
```

**Click here to create your free account in [Developer Sandbox](https://developers.redhat.com/developer-sandbox)**

### Running Locally

```
go run cmd/main.go
```

<br>

## ⛏️ Built Using <a name = "built_using"></a>

- [telegram-bot-api](https://go-telegram-bot-api.dev/) - Golang bindings for the Telegram Bot API
- [kubernetes/client-go]("https://github.com/kubernetes/client-go") - Golang bindings for Kubernetes
- [Openshift](https://www.redhat.com/en/technologies/cloud-computing/openshift) - PaaS hosting platform

<br>

## ✍️ Authors <a name = "authors"></a>

- [@Lucifergene](https://github.com/Lucifergene) - Idea & Implementation

<br>

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- @RedHatOfficial for promoting and organizing hackathons as part of the work culture 🎉
- Inspired by the popular Telegram Bots I use everyday!
