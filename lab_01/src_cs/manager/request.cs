namespace ResourceRegistrator.Manager
{
    public class Request
    {
        public string ResourceId { get; set; }
        public int Time { get; set; }
        public string Status { get; set; }

        public Request(string resourceId, int time)
        {
            ResourceId = resourceId;
            Time = time;
            Status = "Waiting";
        }

        public string GetResourceId()
        {
            return ResourceId;
        }

        public bool IsActive()
        {
            return Status == "Active";
        }

        public bool IsWaiting()
        {
            return Status == "Waiting";
        }

        public void MakeActive()
        {
            Status = "Active";
        }

        public void Process()
        {
            Time -= 1;
        }

        public bool IsOver()
        {
            return Time <= 0;
        }

        public override string ToString()
        {
            return $"\tRequest ResourceId={ResourceId}, Time Left={Time}, Status={Status}";
        }
    }

    public class RequestSlice : List<Request>
        {
            public override string ToString()
            {
                var strs = new List<string>(this.Count);
                foreach (var r in this)
                {
                    strs.Add(r.ToString());
                }
                return $"[\n{string.Join("\n", strs)}\n]";
            }
        }
}
