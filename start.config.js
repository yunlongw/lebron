module.exports = {
    apps: [
        {
            name: 'user',
            script: "go",
            args: ["run", "user.go"],
            interpreter: 'none',
            autorestart: false,
            watch: false,
            cwd: 'E:\\www\\lebron\\apps\\user\\rpc',
        },
        {
            name: 'pay',
            script: "go",
            args: ["run", "pay.go"],
            interpreter: 'none',
            autorestart: false,
            watch: false,
            cwd: 'E:\\www\\lebron\\apps\\pay\\rpc',
        },
        {
            name: 'product',
            script: "go",
            args: ["run", "product.go"],
            interpreter: 'none',
            autorestart: false,
            watch: false,
            cwd: 'E:\\www\\lebron\\apps\\product\\rpc',
        },
        {
            name: 'order',
            script: "go",
            args: ["run", "order.go"],
            interpreter: 'none',
            autorestart: false,
            watch: false,
            cwd: 'E:\\www\\lebron\\apps\\order\\rpc',
        },

    ],
};