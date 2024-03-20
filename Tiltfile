load('ext://min_tilt_version', 'min_tilt_version')
min_tilt_version('0.33.11')

def proxybot():
    docker_build('proxybot', '.',
        dockerfile='proxybot/Dockerfile',
        only=['libs', 'proxybot'],
        target='dev',
        ignore=[
            '*',
            '!proxybot/go.mod',
            '!proxybot/go.sum',
            '!proxybot/cmd',
            '!proxybot/config',
            '!proxybot/pkg',
            '!libs',
        ],
        live_update=[
            sync('proxybot', '/tmp/proxybot'),
            run('go mod download', trigger=['proxybot/go.mod', 'proxybot/go.sum']),
            run('CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /tmp/proxybot.app -gcflags="all=-N -l" cmd/server/main.go'),
            restart_container()
       ]
    )
    dc_resource('proxybot',
        trigger_mode=TRIGGER_MODE_AUTO,
        auto_init=True,
    )

def proxy():
    docker_build('proxy', '.',
        dockerfile='proxy/Dockerfile',
        only=['libs', 'proxy'],
        target='dev',
        ignore=[
            '*',
            '!proxy/go.mod',
            '!proxy/go.sum',
            '!proxy/cmd',
            '!libs',
        ],
        live_update=[
            sync('proxy', '/tmp/proxy'),
            run('go mod download', trigger=['proxy/go.mod', 'proxy/go.sum']),
            run('CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /tmp/proxy.app -gcflags="all=-N -l" cmd/server/main.go'),
            restart_container()
       ]
    )
    dc_resource('proxy',
        trigger_mode=TRIGGER_MODE_AUTO,
        auto_init=True,
    )

def backend():
    docker_build('backend', '.',
        dockerfile='backend/Dockerfile',
        only=['libs', 'backend'],
        target='dev',
        ignore=[
            '*',
            '!backend/go.mod',
            '!backend/go.sum',
            '!backend/cmd',
            '!backend/internal',
            '!backend/pkg',
            '!backend/tools.go',
            '!libs',
        ],
        live_update=[
            sync('backend', '/tmp/backend'),
            run('go mod download', trigger=['backend/go.mod', 'backend/go.sum']),
            run('CGO_ENABLED=0 GOARCH="amd64" GOOS="linux" go build -o /tmp/backend.app -gcflags="all=-N -l" cmd/server/main.go'),
            restart_container()
       ]
    )
    dc_resource('backend',
        trigger_mode=TRIGGER_MODE_AUTO,
        resource_deps=['proxy'],
        auto_init=True,
    )

docker_compose('./docker-compose.yaml')
proxy()
backend()
proxybot()
