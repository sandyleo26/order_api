-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS order_record (
  id integer NOT NULL,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  start_lat numeric,
  start_long numeric,
  end_lat numeric,
  end_long numeric,
  distance numeric,
  status text
);

CREATE SEQUENCE order_record_id_seq
  START with 1
  INCREMENT BY 1
  NO MINVALUE
  NO MAXVALUE
  CACHE 1;

ALTER SEQUENCE order_record_id_seq OWNED BY order_record.id;
ALTER TABLE ONLY order_record ALTER COLUMN id SET DEFAULT nextval('order_record_id_seq'::regclass);
ALTER TABLE ONLY order_record ADD CONSTRAINT order_record_pkey PRIMARY KEY (id);
CREATE INDEX idx_order_record_deleted_at ON order_record USING btree (deleted_at);


-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE order_record;
