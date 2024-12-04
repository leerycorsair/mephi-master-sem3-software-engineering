namespace ResourceRegistrator.Model
{
    public enum ResourceStatus
    {
        Free,
        Occupied
    }

    public class Resource
    {
        public ResourceStatus Status { get; private set; }

        public Resource(ResourceStatus status)
        {
            Status = status;
        }

        public bool IsFree()
        {
            return Status == ResourceStatus.Free;
        }

        public void Free()
        {
            Status = ResourceStatus.Free;
        }

        public void Occupy()
        {
            if (IsFree())
            {
                Status = ResourceStatus.Occupied;
            }
            else
            {
                throw new Exception(Errors.ResourceIsBusy);
            }
        }
    }
}
