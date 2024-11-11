package service

import (
	"lab1/entity"
)

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{
		videos: []entity.Video{},
	}
}

func (svc *videoService) FindAll() []entity.Video {
	return svc.videos
}

func (svc *videoService) Save(video entity.Video) entity.Video {
	svc.videos = append(svc.videos, video)
	return video
}
