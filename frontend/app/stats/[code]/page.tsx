"use client";
import { Card, Spin, Typography } from "antd";
import { use, useEffect, useState } from "react";

const { Title, Text } = Typography;

interface URLStats {
  original_url: string;
  short_code: string;
  click_count: number;
  created_at: string;
  expires_at: string | null;
}

export default function StatsPage({ params }: { params: Promise<{ code: string }> }) {
  const { code } = use(params);
  const [loading, setLoading] = useState(true);
  const [data, setData] = useState<URLStats | null>(null);

  useEffect(() => {
    fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/stats/${code}`)
      .then((res) => res.json())
      .then((data) => {
        setData(data);
        setLoading(false);
      });
  }, [code]); // ← empty array means "run once on page load"

  return (
    <div className="flex flex-col flex-1 items-center justify-center min-h-screen">
      <Card style={{ width: 480 }}>
        <Title level={3} style={{ marginBottom: 24 }}>
          URL Shortener Stats
        </Title>
        <Spin spinning={loading}>
          {!loading && data && (
            <div style={{ display: "flex", flexDirection: "column", gap: 12 }}>
              <Text>
                <b>Original URL:</b> {data.original_url}
              </Text>
              <Text>
                <b>Short Code:</b> {data.short_code}
              </Text>
              <Text>
                <b>Click Count:</b> {data.click_count}
              </Text>
              <Text>
                <b>Created At:</b>{" "}
                {new Date(data.created_at).toLocaleString("en-GB", {
                  day: "numeric",
                  month: "short",
                  year: "numeric",
                  hour: "2-digit",
                  minute: "2-digit",
                  hour12: true,
                })}
              </Text>
              <Text>
                <b>Expires At:</b>{" "}
                {data.expires_at
                  ? new Date(data.expires_at).toLocaleString("en-GB", {
                      day: "numeric",
                      month: "short",
                      year: "numeric",
                      hour: "2-digit",
                      minute: "2-digit",
                      hour12: true,
                    })
                  : "No expiry"}
              </Text>
            </div>
          )}
        </Spin>
      </Card>
    </div>
  );
}
