/**
 * Message Publisher Port
 *
 * Abstraction for publishing messages to a message queue.
 * This allows the core application to be decoupled from the specific
 * messaging implementation (RabbitMQ, Kafka, Redis, etc.)
 */

export interface PublishOptions {
  persistent?: boolean;
  contentType?: string;
  headers?: Record<string, any>;
}

export interface IMessagePublisher {
  /**
   * Publish a message to a queue
   * @param queueName - Name of the queue to publish to
   * @param message - The message to publish (will be JSON serialized)
   * @param options - Optional publishing options
   */
  publish(
    queueName: string,
    message: any,
    options?: PublishOptions
  ): Promise<void>;

  /**
   * Ensure a queue exists before publishing
   * @param queueName - Name of the queue to assert
   */
  assertQueue(queueName: string): Promise<void>;
}
