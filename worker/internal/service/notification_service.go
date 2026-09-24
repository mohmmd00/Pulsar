package service

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/mohmmd00/Pulsar/worker/internal/models"
	"github.com/mohmmd00/Pulsar/worker/internal/repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ProcessNotification(noteEvent models.NotificationEvent, pool *pgxpool.Pool, ctx context.Context) error {

	errMsg := "couldnt process notification : "
	fetchedNote, getNoteErr := repository.GetNotificationByID(noteEvent.NotificationID, pool, ctx)
	if getNoteErr != nil {
		return fmt.Errorf(errMsg+"%w", getNoteErr)
	}

	time.Sleep(2 * time.Second)

	newStatus := "failed"  // high chance to fail to process notification 
	if rand.Intn(10) > 7 { // only 8 and 9
		newStatus = "sent"
	}

	//repository layer
	repoErr := repository.UpdateNotificationStatus(fetchedNote.ID , newStatus , pool , ctx)
	if repoErr != nil {
		return fmt.Errorf(errMsg +"%w" , repoErr)
	}
	return nil
	

}
