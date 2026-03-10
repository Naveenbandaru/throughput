# throughput
**Correlated Telemetry Driven Throughput Analysis in Distributed Pipelines**

### Paper Information
- **Author(s):** Naveen Kumar Bandaru
- **Published In:** Computer Farud and Security
- **Publication Date:** May, 2022
- **ISSN:** 1873-7056
- **DOI:**
- **Impact Factor:** 9.56

### Abstract
Throughput degradation in distributed data pipelines becomes more pronounced as systems scale across multiple nodes and operate under dynamic workloads. Traditional monitoring methods analyze metrics, logs, and traces independently, limiting comprehensive visibility into execution behavior. This study investigates a correlated telemetry approach that integrates these signals using shared execution identifiers and temporal alignment. The analysis examines how execution flow, coordination overhead, and Input Output activity influence throughput across different cluster sizes. The proposed perspective supports a structured understanding of scalability behavior and throughput dynamics in distributed data pipelines.

### Key Contributions
- **Correlated Telemetry Based Analysis:**  
Introduced an approach that connects metrics, logs, and traces through shared execution identifiers and time alignment to examine throughput behavior in distributed data pipelines.

- **Integrated Observability Framework:**  
Established a monitoring structure that combines multiple telemetry signals to study execution flow, coordination overhead, and Input Output interactions across pipeline stages.

- **Distributed Pipeline Experimental Model:** 
Implemented a pipeline simulation that produces telemetry from different processing stages to investigate execution behavior and throughput dynamics in distributed environments.

- **Scalability Evaluation Across Cluster Sizes:**  
Performed experimental analysis on clusters with 3, 5, 7, 9, and 11 nodes to observe how throughput changes as distributed systems scale.

### Practical Significance and Impact
- **Clear Visibility into Throughput Behavior:**
Correlated telemetry enables systematic observation of how execution flow and resource interactions influence throughput in distributed pipelines.

- **Improved Identification of Scalability Constraints:**  
The analysis helps reveal coordination overhead and shared resource contention that affect system performance as cluster size increases.

- **Support for Data Processing Platforms:**  
The framework provides insights useful for improving performance in stream processing systems, analytics platforms, and large scale data pipelines.

- **Foundation for Observability Driven Performance Analysis:**  
Demonstrates the value of integrating telemetry signals to understand distributed execution behavior rather than relying on isolated monitoring approaches.

### Experimental Results (Summary)

  | Nodes | Reference Configuration throughput (ops/sec) | Telemetry Integrated Configuration throughput (ops/sec) | Improvment (%)  |
  |-------|----------------------------------------------| --------------------------------------------------------| ----------------|
  | 3     |  420                                         | 510                                                     | 21.43           |
  | 5     |  560                                         | 680                                                     | 21.43           |
  | 7     |  610                                         | 740                                                     | 21.31           |
  | 9     |  590                                         | 720                                                     | 22.03           |
  | 11    |  540                                         | 690                                                     | 27.78           |

### Citation
Correlated Telemetry Driven Throughput Analysis in Distributed Pipelines
* Naveen Kumar Bandaru
* Computer Farud and Security 
* *****ISSN 1873-7056
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
[**************https://www.ijlrp.com/](https://www.ijirmps.org/research-paper.php?id=232943) \
**Author Contact** \
**LinkedIn**: linkedin.com/in/naveen-bandaru | **Email**: naveen.bandaru@gmail.com







