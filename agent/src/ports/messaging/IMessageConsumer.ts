/**
 * Message Consumer Port
 *
 * Abstraction for consuming messages from a message queue.
 * This allows the core application to be decoupled from the specific
 * messaging implementation (RabbitMQ, Kafka, Redis, etc.)
 */

export interface IMessageConsumer<TMessage = any> {
  /**
   * Start consuming messages from the queue
   * @param queueName - Name of the queue to consume from
   * @param handler - Callback to handle incoming messages
   */
  consume(
    queueName: string,
    handler: (message: TMessage) => Promise<void>
  ): Promise<void>;

  /**
   * Stop consuming messages
   */
  stop(): Promise<void>;
}
