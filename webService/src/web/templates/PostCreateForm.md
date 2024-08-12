# Create a New Post

Заполните форму ниже для создания нового поста.

## Short Description

_Введите краткое описание (до 1000 символов):_

```html
<form action="/create-post" method="POST">
    <label for="shortDescription">Short Description:</label><br>
    <textarea id="shortDescription" name="shortDescription" rows="4" cols="50" maxlength="1000" placeholder="Enter a brief description..."></textarea><br><br>

    <label for="description">Description:</label><br>
    <textarea id="description" name="description" rows="10" cols="50" maxlength="10000" placeholder="Enter the full description..."></textarea><br><br>

    <input type="submit" value="Create Post">
</form>
```