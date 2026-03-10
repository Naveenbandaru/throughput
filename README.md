# throughput
**Correlated Telemetry Driven Throughput Analysis in Distributed Pipelines**

### Paper Information
- **Author(s):** Arunkumar Sambandam
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

### Relevance & Real-World Impact
- **Reduced CPU Utilization:**
Achieved significant processor efficiency by eliminating redundant monitoring tasks, freeing resources for application workloads and improving responsiveness.

- **Improved Latency and Throughput:**  
Lower monitoring overhead directly reduces latency and increases throughput, enhancing service stability in distributed environments.

- **Scalable Cloud Deployment:**  
Framework supports large clusters without proportional CPU growth, addressing scalability challenges in modern cloud infrastructures.

- **Operational Cost and Energy Savings:**  
Efficient monitoring reduces unnecessary computation, lowering energy consumption and operational costs in data centers.

- **Practical Applicability:**
Provides a reference model for industry and research, offering a processor‑efficient monitoring design suitable for production systems and academic exploration.
 
### Experimental Results (Summary)

  | Nodes | Local Telemetry CPU | Telemetry corelation CPU | Improvment (%)  |
  |-------|---------------------| -------------------------| ----------------|
  | 3     |  72                 | 54                       | 25.0           |
  | 5     |  70                 | 50                       | 28.6           |
  | 7     |  68                 | 47                       | 30.9           |
  | 9     |  67                 | 45                       | 32.8           |
  | 11    |  66                 | 43                       | 34.8           |

### Citation
Cross Node Telemetry for CPU Efficient Congestion Monitoring
* Arunkumar Sambandam
* International Journal of Leading Research Publication 
* *****ISSN E-ISSN: 2582-8010
* License \
This research is shared for a academic and research purposes. For commercial use, please contact the author.\
**Resources** \
[**************https://www.ijlrp.com/](https://www.ijirmps.org/research-paper.php?id=232943) \
**Author Contact** \
**LinkedIn**: linkedin.com/in/arunkumar-sambandam-9b769b | **Email**: arunkumar.sambandam@yahoo.com







