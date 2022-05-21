# Go-Flutter Firebase Push Notification

## Usage
1. Create a .env file, put

```bash
FIREBASE_API_KEY=
FIREBASE_DEVICE_TOKEN_ID=
```

* FIREBASE_API_KEY stands for the Server Key of Cloud Messaging API (legacy) in firebase console -> project settings -> cloud messaging
* FIREBASE_DEVICE_TOKEN_ID is acquired from device

```
    _firebaseMessaging.getToken().then((String? token) {
      if (!kReleaseMode) print('fcm token => $token');
    });
```

2. Run the main.go file (go run main.go)
3. Check incoming notification