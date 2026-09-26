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

	//repository layer
	statToProcessingErr := repository.UpdateNotificationStatus(fetchedNote.ID, "processing", false, pool, ctx) //do not increment attempt
	if statToProcessingErr != nil {
		return fmt.Errorf(errMsg+"%w", statToProcessingErr)
	}

	time.Sleep(3 * time.Second)

	newStatus := "failed"  // high chance to fail to process notification
	if rand.Intn(10) > 7 { // only 8 and 9
		newStatus = "sent"
	}

	//repository layer
	statOnChanceErr := repository.UpdateNotificationStatus(fetchedNote.ID, newStatus, true, pool, ctx) //do increment attempt
	if statOnChanceErr != nil {
		return fmt.Errorf(errMsg+"%w", statOnChanceErr)
	}

	var result string
	var errMessage string

	if newStatus == "failed" {
		result = "failure"
		errMessage = "failed to process notification"
	} else {
		result = "success"
		errMessage = ""
	}

	noteLogErr := repository.CreateNotificationLog(models.NewNotificationLog(fetchedNote.ID, result, errMessage), pool, ctx)

	if noteLogErr != nil {
		fmt.Println("failed to write notification log:", noteLogErr)
	}

	return nil

}
