CREATE TABLE IF NOT EXISTS log_analysis(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    analysis text,
    cause text,
    severity text,
    time_of_occurrence timestamp(0) with time zone DEFAULT now(),
    stack_trace text,
    file text,
    line_number text,
    summary text,
    comprehensive_details text,
    suggested_way_to_fix text,
    created_at timestamp(0) with time zone DEFAULT now()
);