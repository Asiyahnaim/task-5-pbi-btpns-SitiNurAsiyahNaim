func AddPhoto(c *gin.Context) {
	var newPhoto models.Photo
	if err := c.BindJSON(&newPhoto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Create(&newPhoto)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"photo": newPhoto})
}

func GetAllPhotos(c *gin.Context) {
	var photos []models.Photo
	result := database.DB.Find(&photos)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"photos": photos})
}

func UpdatePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	var updateInfo models.Photo
	if err := c.BindJSON(&updateInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var photo models.Photo
	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	database.DB.Model(&photo).Updates(updateInfo)
	c.JSON(http.StatusOK, gin.H{"photo": photo})
}

func DeletePhoto(c *gin.Context) {
	photoID := c.Param("photoId")

	var photo models.Photo
	result := database.DB.First(&photo, photoID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "photo not found"})
		return
	}

	database.DB.Delete(&photo)
	c.JSON(http.StatusOK, gin.H{"message": "Photo deleted"})
}
