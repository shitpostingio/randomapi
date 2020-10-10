# Memes API

**Get a random meme**
----
  Get a random meme from our collection

* **URL**

  /random

* **Method:**

  `GET`
  
*  **URL Params**

   **Optionals:**
 
   - `type=[image, video, animation]`
   - `startDate=unixdate` get a meme that has been posted after given date
   - `endDate=unixdate` get a meme that has been posted before given date

* **HTTP Headers**
    
    **Required:**
    - `X-user-platform`
    - `X-user-id`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:**
    ``` 
    {
    "id": "bo7od5qka8toqdhaph8g",
    "data": {
        "url": "https://api.shitposting.io/storage/rand_bo7od5qka8toqdhaph8g.jpg",
        "caption": "Is this true? We need more straight men, and women, in @thememaly RIGHT NOW\n\nJoin us and discuss\n\n\\[By BlueTag]",
        "filename": "rand_bo7od5qka8toqdhaph8g.jpg",
        "messageid": 18048,
        "mediatype": "image",
        "date": "2018-03-30T10:55:09+02:00"
        }
    }
    ```
