package data

import (
	db "github.com/hi20160616/youtube_web/internal/pkg/db/json"
	"google.golang.org/api/youtube/v3"
)

// type ChannelsRepo struct {
//         Channels   []*db.Channel
//         VideoLists []*youtube.VideoListResponse
//         // PublishAfter int // how many minutes ago
// }

// func (cr *ChannelsRepo) ReadChannels() (*ChannelsRepo, error) {
//         cs, err := db.ReadChannels()
//         if err != nil {
//                 return nil, err
//         }
//         cr.Channels = cs
//         return cr, nil
// }
// TODO: mv to db layer
// func (cr *ChannelsRepo) GetAllActivitiesVideos() (*ChannelsRepo, error) {
//         for _, c := range cr.Channels {
//                 vr, err := NewVideosRepo().
//                         WithChannelId(c.ChannelId).
//                         WithPublishAfter4Activities(cr.PublishAfter).
//                         GetActivitiesVideos()
//                 if err != nil {
//                         if errors.Is(err, db.ErrAcitvityListNil) {
//                                 continue
//                         }
//                         return nil, err
//                 }
//                 cr.VideoLists = append(cr.VideoLists, vr)
//         }
//         return cr, nil
// }
//
// func (cr *ChannelsRepo) WithPublishAfter(minutes int) *ChannelsRepo {
//         cr.PublishAfter = minutes
//         return cr
// }

func ReadChannels() ([]*db.Channel, error) {
	return db.ReadChannels()
}

func ReadActivities() ([]*youtube.VideoListResponse, error) {
	return db.ReadActivities()
}
