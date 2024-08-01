const net = require('net');

// Função para criar e iniciar um servidor em uma porta específica
function createServer(port) {
    const server = net.createServer((socket) => {
        console.log(`Cliente conectado na porta ${port}`);

        socket.setEncoding('utf8');

        socket.on('data', (data) => {
            console.log(`Recebido na porta ${port}: ${data}`);
            socket.write(`Reposta ${port}: ${data}`);
        });

        socket.on('end', () => {
            console.log(`Cliente desconectado da porta ${port}`);
        });

        socket.on('error', (err) => {
            console.error(`Erro na porta ${port}:`, err);
        });
    });

    server.listen(port, () => {
        console.log(`Servidor TCP escutando na porta ${port}`);
    });
}

// Cria servidores nas portas de 10000 a 11000
for (let port = 10000; port <= 11000; port++) {
    createServer(port);
}
