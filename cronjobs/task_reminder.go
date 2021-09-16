package cronjobs

import (
	"fmt"
	"github.com/ogbofjnr/maze/database/repositories"
	"github.com/ogbofjnr/maze/notifications"
	"github.com/ogbofjnr/maze/pkg/db"
	"github.com/ogbofjnr/maze/pkg/mailer"
	"github.com/ogbofjnr/maze/pkg/mailer/emails"
	"github.com/ogbofjnr/maze/pkg/notificator"
	"go.uber.org/zap"
)

type TaskReminder struct {
	db             db.DB
	userRepository *repositories.UserRepository
	taskRepository *repositories.TaskRepository
	logger         *zap.Logger
	notificator    *notificator.Notificator
	mailer         *mailer.Mailer
}

func NewTaskReminder(
	db db.DB,
	userRepository *repositories.UserRepository,
	TaskRepository *repositories.TaskRepository,
	logger *zap.Logger,
	notificator *notificator.Notificator,
	mailer *mailer.Mailer,
) *TaskReminder {
	t := &TaskReminder{
		db:             db,
		userRepository: userRepository,
		taskRepository: TaskRepository,
		logger:         logger,
		notificator:    notificator,
		mailer:         mailer,
	}
	return t
}

// Run contains logic of the cron task
func (s *TaskReminder) Run() error {

	tasks, err := s.taskRepository.GetTasksForNotification(s.db)
	if err != nil {
		s.logger.Error(fmt.Sprintf("error in cron job %s", err.Error()))
		return err
	}
	for _, task := range *tasks {
		user, err := s.userRepository.GetByID(s.db, task.ID)
		if err != nil {
			s.logger.Error(fmt.Sprintf("error in cron job %s", err.Error()))
			return err
		}
		notification := notifications.NewTaskReminderNotification(task.Title, task.UUID)
		s.notificator.Notify(user.ID, notification)
		err = s.mailer.Send(emails.NewTaskReminder(user.Email, user.GetFullName(), ""))
		if err != nil {
			s.logger.Error(fmt.Sprintf("error in cron job %s", err.Error()))
			return err
		}
	}

	return nil
}
