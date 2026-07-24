package database

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/cb00j/cbj-rwa/rwa-backend/libs/log"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

type Migration struct {
	client *migrate.Migrate
}

// NewMigration todo log should ignore password
func NewMigration(ctx context.Context, conf *DbConf) (*Migration, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s&x-migrations-table=%s",
		conf.Username,
		conf.Password,
		net.JoinHostPort(conf.Host, strconv.Itoa(conf.Port)),
		conf.Database,
		conf.SslMode,
		"_migrations",
	)

	// 不再用 file:// 这种URL字符串——golang-migrate 的 file:// 驱动在Windows上
	// 对路径的解析有实测确认过的缺陷(不管是相对路径、还是带///的绝对路径,
	// url.Parse 解析出来的 Path 字段要么被截断、要么带着Windows无法识别的
	// 前导斜杠)。改用基于 io/fs 的 iofs 驱动 + os.DirFS,直接用Go标准库原生
	// 处理文件系统路径,完全不经过URL字符串解析这一步,Windows/macOS通用。
	relPath := strings.TrimPrefix(conf.MigrationPath, "file://")
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		log.ErrorZ(ctx, "failed to resolve migration path", zap.Error(err), zap.String("configuredPath", conf.MigrationPath))
		return nil, err
	}
	log.InfoZ(ctx, "resolved migration directory", zap.String("configuredPath", conf.MigrationPath), zap.String("resolvedAbsPath", absPath))

	fsys := os.DirFS(absPath)
	sourceDriver, err := iofs.New(fsys, ".")
	if err != nil {
		log.ErrorZ(ctx, "failed to create iofs source driver", zap.Error(err), zap.String("absPath", absPath))
		return nil, err
	}

	client, err := migrate.NewWithSourceInstance("iofs", sourceDriver, dbURL)
	if err != nil {
		log.ErrorZ(ctx, "new migration failed", zap.Error(err))
		return nil, err
	}
	return &Migration{
		client: client,
	}, nil
}

func (m *Migration) To(ctx context.Context, targetVersion uint) error {
	if targetVersion == 0 {
		log.WarnZ(ctx, "target version is 0")
		return nil
	}
	if err := m.client.Migrate(targetVersion); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.ErrorZ(ctx, "Migration to version failed", zap.Error(err), zap.Uint("version", targetVersion))
		return err
	}
	afterVersion, _, _ := m.client.Version()
	log.InfoZ(ctx, "Migration to version success", zap.Uint("version", afterVersion))
	return nil
}

func (m *Migration) Up(ctx context.Context) error {
	if err := m.client.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.ErrorZ(ctx, "Migration up version failed", zap.Error(err))
		return err
	}
	afterVersion, _, _ := m.client.Version()
	log.InfoZ(ctx, "Migration up version success", zap.Uint("version", afterVersion))
	return nil
}

func (m *Migration) Down(ctx context.Context) error {
	if err := m.client.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.ErrorZ(ctx, "Migration down version failed", zap.Error(err))
		return err
	}
	version, _, _ := m.client.Version()
	log.InfoZ(ctx, "Migration down version success", zap.Uint("version", version))
	return nil
}
