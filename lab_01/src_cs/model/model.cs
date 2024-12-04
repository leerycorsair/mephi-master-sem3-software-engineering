using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace ResourceRegistrator.Model
{
    public class Model
    {
        private Dictionary<string, Resource> resources = new Dictionary<string, Resource>();

        public void AddResource(string nameTag, ResourceStatus status)
        {
            var resource = new Resource(status);
            resources[nameTag] = resource;
        }

        public void FreeResource(string nameTag)
        {
            if (!resources.TryGetValue(nameTag, out var resource))
            {
                throw new Exception(Errors.ResourceNotFound);
            }
            resource.Free();
            resources[nameTag] = resource;
        }

        public Resource SearchResource(string nameTag)
        {
            if (resources.TryGetValue(nameTag, out var resource))
            {
                return resource;
            }
            return null;
        }

        public void OccupyResource(string nameTag)
        {
            if (!resources.TryGetValue(nameTag, out var resource))
            {
                throw new Exception(Errors.ResourceNotFound);
            }
            resource.Occupy();
            resources[nameTag] = resource;
        }

        public void InitResources(int resourceCount)
        {
            for (int i = 0; i < resourceCount; i++)
            {
                AddResource($"res{i}", ResourceStatus.Free);
            }
        }

        public override string ToString()
        {
            var sb = new StringBuilder();
            foreach (var resource in resources)
            {
                sb.AppendLine($"{resource.Key}: {resource.Value.Status}");
            }
            return sb.ToString();
        }

        public bool IsEmpty()
        {
            return resources.Count == 0;
        }
    }
}
