/**
 * RabbitMQ Adapter
 *
 * Concrete implementation of messaging ports using RabbitMQ.
 * Handles connection management, queue assertion, and message operations.
 */

import amqp from "amqplib";
import type { IMessageConsumer } from "../../ports/messaging/IMessageConsumer.js";
import type {
  IMessagePublisher,
  PublishOptions,
} from "../../ports/messaging/IMessagePublisher.js";

export interface RabbitMQConfig {
  url: string;
  onConnectionError?: (error: Error) => void;
  onConnectionClose?: () => void;
}

export class RabbitMQAdapter implements IMessageConsumer, IMessagePublisher {
  private connectionModel: amqp.ChannelModel | null = null;
  private channel: amqp.Channel | null = null;
  private readonly config: RabbitMQConfig;
  private consumerTag: string | null = null;

  constructor(config: RabbitMQConfig) {
    this.config = config;
  }

  /**
   * Connect to RabbitMQ and create a channel
   */
  async connect(): Promise<void> {
    try {
      this.connectionModel = await amqp.connect(this.config.url);
      this.channel = await this.connectionModel.createChannel();

      // Handle connection errors
      this.connectionModel.on("error", (err: Error) => {
        console.error("[RabbitMQ] Connection error:", err);
        if (this.config.onConnectionError) {
          this.config.onConnectionError(err);
        }
      });

      this.connectionModel.on("close", () => {
        console.log("[RabbitMQ] Connection closed");
        if (this.config.onConnectionClose) {
          this.config.onConnectionClose();
        }
      });

      console.log("[RabbitMQ] Connected successfully");
    } catch (error) {
      console.error("[RabbitMQ] Failed to connect:", error);
      throw error;
    }
  }

  /**
   * Assert that a queue exists (create if it doesn't)
   */
  async assertQueue(queueName: string): Promise<void> {
    if (!this.channel) {
      throw new Error("Channel not initialized. Call connect() first.");
    }

    try {
      await this.channel.assertQueue(queueName, { durable: true });
    } catch (error) {
      console.error(`[RabbitMQ] Failed to assert queue ${queueName}:`, error);
      throw error;
    }
  }

  /**
   * Publish a message to a queue
   */
  async publish(
    queueName: string,
    message: any,
    options?: PublishOptions
  ): Promise<void> {
    if (!this.channel) {
      throw new Error("Channel not initialized. Call connect() first.");
    }

    try {
      const content = Buffer.from(JSON.stringify(message));
      const publishOptions: amqp.Options.Publish = {
        contentType: options?.contentType || "application/json",
        persistent: options?.persistent ?? true,
        headers: options?.headers,
      };

      this.channel.sendToQueue(queueName, content, publishOptions);
    } catch (error) {
      console.error(`[RabbitMQ] Failed to publish to ${queueName}:`, error);
      throw error;
    }
  }

  /**
   * Start consuming messages from a queue
   */
  async consume<TMessage = any>(
    queueName: string,
    handler: (message: TMessage) => Promise<void>
  ): Promise<void> {
    if (!this.channel) {
      throw new Error("Channel not initialized. Call connect() first.");
    }

    try {
      const consumeResult = await this.channel.consume(
        queueName,
        async (msg) => {
          if (!msg) {
            return;
          }

          try {
            const message: TMessage = JSON.parse(msg.content.toString());
            await handler(message);

            // Acknowledge message after successful processing
            this.channel?.ack(msg);
          } catch (error) {
            console.error("[RabbitMQ] Error processing message:", error);

            // Acknowledge anyway to prevent requeue loop
            // In production, you might want to:
            // - Requeue with a retry limit
            // - Send to a dead letter queue
            // - Use nack with requeue: false
            this.channel?.ack(msg);
          }
        },
        { noAck: false } // Manual acknowledgment
      );

      this.consumerTag = consumeResult.consumerTag;
      console.log(`[RabbitMQ] Started consuming from ${queueName}`);
    } catch (error) {
      console.error(`[RabbitMQ] Failed to consume from ${queueName}:`, error);
      throw error;
    }
  }

  /**
   * Stop consuming messages
   */
  async stop(): Promise<void> {
    try {
      if (this.consumerTag && this.channel) {
        await this.channel.cancel(this.consumerTag);
        this.consumerTag = null;
      }
    } catch (error) {
      console.error("[RabbitMQ] Error stopping consumer:", error);
    }
  }

  /**
   * Disconnect from RabbitMQ
   */
  async disconnect(): Promise<void> {
    try {
      await this.stop();

      if (this.channel) {
        await this.channel.close();
        this.channel = null;
      }

      if (this.connectionModel) {
        await this.connectionModel.close();
        this.connectionModel = null;
      }

      console.log("[RabbitMQ] Disconnected");
    } catch (error) {
      console.error("[RabbitMQ] Error during disconnect:", error);
    }
  }

  /**
   * Check if connected
   */
  isConnected(): boolean {
    return this.connectionModel !== null && this.channel !== null;
  }
}
