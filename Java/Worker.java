import java.util.List;
import java.util.logging.Logger;

public class Worker extends Thread {
    private final TaskQueue taskQueue;
    private final List<String> results;
    private final Logger logger;
    private final String workerName;

    public Worker(TaskQueue taskQueue, List<String> results, Logger logger, String name) {
        this.taskQueue = taskQueue;
        this.results = results;
        this.logger = logger;
        this.workerName = name;
    }

    @Override
    public void run() {
        logger.info(workerName + " started.");
        while (true) {
            String task = taskQueue.getTask();
            if (task == null) break;

            try {
                Thread.sleep((long) (Math.random() * 1000)); // Simulate processing
                synchronized (results) {
                    results.add(workerName + " processed: " + task);
                }
                logger.info(workerName + " completed: " + task);
            } catch (InterruptedException e) {
                logger.severe(workerName + " interrupted: " + e.getMessage());
            }
        }
        logger.info(workerName + " finished all tasks.");
    }
}
