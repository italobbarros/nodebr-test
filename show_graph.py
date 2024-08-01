import os
import json
import matplotlib.pyplot as plt
import pandas as pd

def read_results(directory):
    results = {}
    for filename in os.listdir(directory):
        if filename.endswith(".json"):
            port = int(filename.split('cliente')[1].split('.json')[0])
            filepath = os.path.join(directory, filename)
            with open(filepath, 'r') as file:
                data = json.load(file)
                results[port] = data['average_response_time']
    return results

def plot_results(go_results, node_results):
    ports = sorted(go_results.keys())
    go_times = [go_results[port] for port in ports]
    node_times = [node_results.get(port, float('inf')) for port in ports]
    

    plt.figure(figsize=(12, 6))

    plt.plot(ports, go_times, label='Go', marker='o')
    plt.plot(ports, node_times, label='Node', marker='o')
    #plt.plot(ports, node_optimized_times, label='Node-optimized', marker='o')

    plt.xlabel('Porta (Cliente)')
    plt.ylabel('Tempo Médio de Resposta (segundos)')
    plt.title('Comparação de Tempo Médio de Resposta')
    plt.legend()
    plt.grid(True)
    plt.savefig('response_times_comparison.png')
    plt.show()

def main():
    # Diretórios dos resultados
    go_dir = 'result/go'
    node_dir = 'result/node'
    # Ler resultados
    go_results = read_results(go_dir)
    node_results = read_results(node_dir)
    # Criar gráfico
    plot_results(go_results, node_results)

if __name__ == "__main__":
    main()
