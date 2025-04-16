import java.io.FileWriter;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;
import java.util.logging.Logger;

public class Main {
    public static void main(String[] args) {
        Logger logger = Logger.getLogger("DataProcessing");
        TaskQueue taskQueue = new TaskQueue();
        List<String> results = new ArrayList<>();

        // Add tasks
        for (int i = 1; i <= 10; i++) {
            taskQueue.addTask("Task " + i);
        }

        // Create and start workers
        List<Worker> workers = new ArrayList<>();
        for (int i = 0; i < 4; i++) {
            Worker worker = new Worker(taskQueue, results, logger, "Worker-" + (i + 1));
            workers.add(worker);
            worker.start();
        }

        // Wait for all workers
        for (Worker worker : workers) {
            try {
                worker.join();
            } catch (InterruptedException e) {
                logger.severe("Join interrupted: " + e.getMessage());
            }
        }

        // Write results
        try (FileWriter writer = new FileWriter("results.txt")) {
            for (String res : results) {
                writer.write(res + "\n");
            }
        } catch (IOException e) {
            logger.severe("File write failed: " + e.getMessage());
        }

        logger.info("All tasks processed and results saved.");
    }
}
