/**
 * Agent Service Bootstrap
 *
 * Thin entry point that wires up dependencies and starts the agent.
 * All business logic is delegated to the application layer.
 */

import { applicationConfig } from 'configuration.js';
import { RabbitMQAdapter, HeartbeatMonitor } from './src/adapters/index.js';
import { AgentService, AgentLifecycle } from './src/application/index.js';

const AGENT_VERSION = '1.0.0';

/**
 * Bootstrap the agent service
 */
(async () => {
  try {
    // Initialize RabbitMQ adapter
    const rabbitMQ = new RabbitMQAdapter({
      url: applicationConfig.mqQueueUrl,
      onConnectionError: (error) => {
        console.error('[Bootstrap] RabbitMQ connection error:', error);
        process.exit(1);
      },
      onConnectionClose: () => {
        console.log('[Bootstrap] RabbitMQ connection closed');
        process.exit(1);
      }
    });

    // Connect to RabbitMQ
    await rabbitMQ.connect();

    // Assert required queues
    await rabbitMQ.assertQueue(applicationConfig.mqHertbeatUrl);
    await rabbitMQ.assertQueue(applicationConfig.mqRequestUrl);

    // Initialize heartbeat monitor
    const heartbeatMonitor = new HeartbeatMonitor({
      queueName: applicationConfig.mqHertbeatUrl,
      publisher: rabbitMQ
    });

    // Initialize agent service
    const agentService = new AgentService({
      version: AGENT_VERSION,
      browserType: 'chromium',
      browserPath: applicationConfig.browserPath,
      messageConsumer: rabbitMQ,
      messagePublisher: rabbitMQ,
      healthMonitor: heartbeatMonitor,
      requestQueueName: applicationConfig.mqRequestUrl
    });

    // Initialize lifecycle manager
    const lifecycle = new AgentLifecycle({
      agentService,
      healthMonitor: heartbeatMonitor,
      onShutdown: async () => {
        await rabbitMQ.disconnect();
      }
    });

    // Start the agent
    await lifecycle.start();
  } catch (error) {
    console.error('[Bootstrap] Fatal error:', error);
    process.exit(1);
  }
})();
