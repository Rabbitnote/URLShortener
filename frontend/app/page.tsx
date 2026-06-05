"use client";
import { Button, Card, DatePicker, Form, Input, Spin, Typography } from "antd";
import { useState } from "react";

const { Title, Text } = Typography;

export default function Home() {
  const [shortUrl, setShortUrl] = useState("");
  const [loading, setLoading] = useState(false);
  const [form] = Form.useForm();

  return (
    <div className="flex flex-col flex-1 items-center justify-center min-h-screen">
      <Card style={{ width: 480 }}>
        <Title level={3} style={{ marginBottom: 24 }}>URL Shortener</Title>
        <Spin spinning={loading}>
          <Form
            form={form}
            layout="vertical"
            onFinish={async (values) => {
              setLoading(true);
              const res = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/shorten`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                  original_url: values.url,
                  expires_at: values.expires_at
                    ? values.expires_at.toISOString()
                    : null,
                }),
              });
              const data = await res.json();
              setShortUrl(data.short_url);
              setLoading(false);
            }}
          >
            <Form.Item name="url" label="Long URL">
              <Input placeholder="https://example.com" />
            </Form.Item>
            <Form.Item name="expires_at" label="Expire Date (Optional)">
              <DatePicker style={{ width: "100%" }} />
            </Form.Item>
            <Form.Item>
              <Button htmlType="submit" type="primary" loading={loading} block>
                Generate Short URL
              </Button>
            </Form.Item>
          </Form>
          {shortUrl && (
            <div style={{ marginTop: 8 }}>
              <Text type="secondary">Your short URL:</Text>
              <div style={{ display: "flex", alignItems: "center", gap: 8, marginTop: 8 }}>
                <a
                  href={`${process.env.NEXT_PUBLIC_API_URL}/${shortUrl}`}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  {process.env.NEXT_PUBLIC_API_URL}/{shortUrl}
                </a>
                <Button
                  size="small"
                  onClick={() => {
                    navigator.clipboard.writeText(`${process.env.NEXT_PUBLIC_API_URL}/${shortUrl}`);
                  }}
                >
                  Copy
                </Button>
              </div>
            </div>
          )}
        </Spin>
      </Card>
    </div>
  );
}
