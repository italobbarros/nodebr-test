const net = require('net');

async function createServer(port) {
    return new Promise((resolve, reject) => {
        const server = net.createServer((socket) => {
            socket.setEncoding('utf8');

            socket.on('data', (data) => {
                console.log(`Recebido na porta ${port}: ${data}`);
                socket.write(`Resposta ${port}: ${data}`);
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
            resolve();
        });

        server.on('error', (err) => {
            reject(`Erro ao iniciar o servidor na porta ${port}: ${err}`);
        });
    });
}

async function startServers(startPort, endPort, batchSize) {
    for (let port = startPort; port <= endPort; port += batchSize) {
        // Cria servidores em lotes para evitar sobrecarregar o sistema
        const batchPromises = [];
        for (let p = port; p < port + batchSize && p <= endPort; p++) {
            batchPromises.push(createServer(p));
        }
        await Promise.all(batchPromises);
    }
}

(async () => {
    try {
        // Ajuste o tamanho do lote conforme a capacidade do sistema
        const batchSize = 100;
        await startServers(10000, 11000, batchSize);
    } catch (error) {
        console.error('Erro ao iniciar servidores:', error);
    }
})();
